package service

import (
	"context"
	docker "github.com/fsouza/go-dockerclient"
	agentv1 "nodeagent/apis/agent/v1"
	agentconf "nodeagent/controllers/config"
)

const (
	DockerCommandCommit        = "docker.commit"
	DockerCommandCommitAndPush = "docker.commitAndPush"
)

type DockerService struct {
	client *docker.Client
	config *agentconf.Config
}

type DockerPushCommand struct {
	Name string
	Tag  string
}

func NewDockerService(config *agentconf.Config) ActionService {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
	//log:    log.NewHelper("dockerService", logger),
	//conf:   conf,
	return &DockerService{
		config: config,
		client: client,
	}
}

func (s *DockerService) commit(ctx context.Context, req *agentv1.DockerCommitCommand) (string, error) {
	opt := docker.CommitContainerOptions{
		Context:    ctx,
		Container:  req.Container,
		Repository: req.Repository,
		Tag:        req.Tag,
		Author:     req.Author,
		Message:    req.Message,
		Changes:    req.Changes,
	}
	image, err := s.client.CommitContainer(opt)
	if err != nil {
		return "", err
	}

	return image.ID, nil
}

func (s *DockerService) push(ctx context.Context, req *DockerPushCommand) error {
	registryHost := s.config.Harbor.Host
	opt1 := docker.PushImageOptions{
		Context:  ctx,
		Name:     req.Name,
		Registry: registryHost,
		Tag:      req.Tag,
	}

	serverAddress := registryHost
	if s.config.Harbor.SSL {
		serverAddress = "https://" + serverAddress
	} else {
		serverAddress = "http://" + serverAddress
	}
	optAuth := docker.AuthConfiguration{
		Username:      s.config.Harbor.Username,
		Password:      s.config.Harbor.Password,
		ServerAddress: serverAddress,
	}
	err := s.client.PushImage(opt1, optAuth)
	if err != nil {
		return err
	}
	return nil
}

func (s *DockerService) commitAndPush(ctx context.Context, req *agentv1.DockerCommitCommand) error {
	if _, err := s.commit(ctx, req); err != nil {
		return err
	}

	preq := &DockerPushCommand{
		Name: req.Repository,
		Tag:  req.Tag,
	}
	if err := s.push(ctx, preq); err != nil {
		return err
	}
	return nil
}

func (s *DockerService) Do(ctx context.Context, action agentv1.Action) []*ActionResult {
	dockerAction := action.Docker
	result := []*ActionResult{}

	if dockerAction.Commit != nil {
		CommitResult := &ActionResult{Name: DockerCommandCommit}
		if _, err := s.commit(ctx, dockerAction.Commit); err != nil {
			CommitResult.Error = err
		}
		result = append(result, CommitResult)
	} else if dockerAction.CommitAndPush != nil {
		CommitResultAndPush := &ActionResult{Name: DockerCommandCommitAndPush}
		if err := s.commitAndPush(ctx, dockerAction.CommitAndPush); err != nil {
			CommitResultAndPush.Error = err
		}
		result = append(result, CommitResultAndPush)
	}
	return result
}
