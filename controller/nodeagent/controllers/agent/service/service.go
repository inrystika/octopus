package service

import (
	"context"
	agentv1 "nodeagent/apis/agent/v1"
	agentconf "nodeagent/controllers/config"
)

type ActionResult struct {
	Name  string
	Error error
}

type ActionService interface {
	Do(ctx context.Context, action agentv1.Action) []*ActionResult
}

type ServiceManager struct {
	services []ActionService
}

func (sm *ServiceManager) Do(ctx context.Context, action agentv1.Action) []*ActionResult {
	result := []*ActionResult{}
	for _, svc := range sm.services {
		result = append(result, svc.Do(ctx, action)...)
	}
	return result
}

func NewActionServiceManager(config *agentconf.Config) ActionService {
	dockerService := NewDockerService(config)

	services := []ActionService{}
	services = append(services, dockerService)
	sm := ServiceManager{
		services: services,
	}
	return &sm
}
