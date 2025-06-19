package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type EnvironmentEndpoint struct {
	useCase usecase.EnvironmentUseCase
}

func NewEnvironmentEndpoint(useCase usecase.EnvironmentUseCase) *EnvironmentEndpoint {
	return &EnvironmentEndpoint{useCase: useCase}
}

func (ee *EnvironmentEndpoint) GetClusters() ([]model.ClusterInfo, error) {
	return ee.useCase.ListAvailableClusters()
}

func (ee *EnvironmentEndpoint) GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error) {
	return ee.useCase.GetCurrentEnvironment(env, name)
}

func (ee *EnvironmentEndpoint) GetObjectsView() (model.ObjectMapDto, error) {
	return ee.useCase.GetObjectsView()
}
