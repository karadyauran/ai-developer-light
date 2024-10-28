package service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/model"
)

type DockerService struct {
	client *client.Client
}

func NewDockerService() *DockerService {
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return &DockerService{client: cli}
}

func (ds *DockerService) CreateContainer(containerConfig model.DockerContainer) (string, error) {
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

	return resp.ID, nil
}
