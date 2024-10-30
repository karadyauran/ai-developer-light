package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"io"
	"path/filepath"
	"time"
)

type DockerService struct {
	client *client.Client
}

func NewDockerService() *DockerService {
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return &DockerService{client: cli}
}

func (ds *DockerService) StartContainer(dockerfilePath string) (string, error) {
	ctx := context.Background()

	absDockerfilePath, err := filepath.Abs(dockerfilePath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	buildContextDir := filepath.Dir(absDockerfilePath)
	buildContext, err := archive.TarWithOptions(buildContextDir, &archive.TarOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to create build context: %w", err)
	}

	imageTag := fmt.Sprintf("myimage:%d", time.Now().Unix())

	buildOptions := types.ImageBuildOptions{
		Dockerfile: filepath.Base(absDockerfilePath),
		Tags:       []string{imageTag},
		Remove:     true,
	}

	buildResponse, err := ds.client.ImageBuild(ctx, buildContext, buildOptions)
	if err != nil {
		return "", fmt.Errorf("failed to build image: %w", err)
	}
	defer buildResponse.Body.Close()

	type buildResponseLine struct {
		Stream string `json:"stream"`
		Error  string `json:"errorDetail"`
	}

	decoder := json.NewDecoder(buildResponse.Body)
	for {
		var line buildResponseLine
		if err := decoder.Decode(&line); err != nil {
			if err == io.EOF {
				break
			}
			return "", fmt.Errorf("error decoding build output: %w", err)
		}
		if line.Stream != "" {
			fmt.Print(line.Stream)
		}
		if line.Error != "" {
			return "", fmt.Errorf("error during image build: %s", line.Error)
		}
	}

	images, err := ds.client.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to list images: %w", err)
	}

	found := false
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == imageTag {
				found = true
				break
			}
		}
	}
	if !found {
		return "", fmt.Errorf("image '%s' not found after build", imageTag)
	}

	resp, err := ds.client.ContainerCreate(ctx, &container.Config{
		Image: imageTag,
		Tty:   true,
	}, nil, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("failed to create container: %w", err)
	}

	containerID := resp.ID

	if err := ds.client.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("failed to start container: %w", err)
	}

	return containerID, nil
}
