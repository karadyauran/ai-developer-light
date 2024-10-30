package service

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/model"
	"os"
	"path/filepath"
)

type DockerService struct {
	client *client.Client
}

func NewDockerService() *DockerService {
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return &DockerService{client: cli}
}

func (ds *DockerService) CopyAndRun(containerConfig model.DockerContainer) (string, error) {
	ctx := context.Background()

	resp, err := ds.client.ContainerCreate(ctx, &container.Config{
		Image:      containerConfig.Image,
		Hostname:   containerConfig.Hostname,
		Tty:        true,
		WorkingDir: "/app",
	}, &container.HostConfig{
		Resources: container.Resources{
			Memory:   512 * 1024 * 1024,
			NanoCPUs: 500000000,
		},
	}, nil, nil, containerConfig.Hostname)
	if err != nil {
		return "", fmt.Errorf("cannot create container:%w", err)
	}
	containerID := resp.ID

	for _, file := range containerConfig.Files {
		if err := ds.copyFileToContainer(containerID, file.Filename, "app/"); err != nil {
			return "", fmt.Errorf("failed to copy file %s: %w", file, err)
		}
	}

	if err := ds.client.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("cannot start container: %w", err)
	}

	for _, cmd := range containerConfig.Commands {
		if err := ds.execInContainer(containerID, cmd); err != nil {
			return "", fmt.Errorf("command %v failed: %w", cmd, err)
		}
	}

	return containerID, nil
}

func (ds *DockerService) copyFileToContainer(containerID, srcPath, destPath string) error {
	ctx := context.Background()
	tarBuffer, err := createTar(srcPath)
	if err != nil {
		return err
	}
	return ds.client.CopyToContainer(ctx, containerID, destPath, tarBuffer, container.CopyToContainerOptions{})
}

func (ds *DockerService) execInContainer(containerID, cmd string) error {
	ctx := context.Background()
	execConfig := container.ExecOptions{
		Cmd:          []string{"/bin/sh", "-c", cmd},
		AttachStdout: true,
		AttachStderr: true,
	}

	execIDResp, err := ds.client.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		return err
	}

	resp, err := ds.client.ContainerExecAttach(ctx, execIDResp.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}
	defer resp.Close()

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, resp.Reader)
	return err
}

func createTar(srcPath string) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	file, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	header, err := tar.FileInfoHeader(fi, fi.Name())
	if err != nil {
		return nil, err
	}
	header.Name = filepath.Base(srcPath)

	if err := tw.WriteHeader(header); err != nil {
		return nil, err
	}
	if _, err := io.Copy(tw, file); err != nil {
		return nil, err
	}

	return buf, nil
}
