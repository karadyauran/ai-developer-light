package service

type Service struct {
	DockerService *DockerService
}

func NewService() *Service {
	return &Service{
		DockerService: NewDockerService(),
	}
}
