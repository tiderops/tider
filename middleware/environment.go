package middleware

import (
	"Kubexplorer/backend/endpoint"
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type EnvironmentMiddleware struct {
	endpoint endpoint.EnvironmentEndpoint
}

func NewEnvironmentMiddleware(endpoint *endpoint.EnvironmentEndpoint) *EnvironmentMiddleware {
	return &EnvironmentMiddleware{endpoint: *endpoint}
}

func (e *EnvironmentMiddleware) GetClusters() ([]model.ClusterInfo, error) {
	return e.endpoint.GetClusters()
}

func (e *EnvironmentMiddleware) GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error) {
	return e.endpoint.GetCurrentEnvironment(env, name)
}

func (e *EnvironmentMiddleware) GetObjectsView() (model.ObjectMapDto, error) {
	return e.endpoint.GetObjectsView()
}

func BuildEnvironment() *EnvironmentMiddleware {
	clusterClient := kubeclient.NewCluster()

	environmentUseCase := usecase.NewEnvironmentUseCase(clusterClient)

	environmentEndpoint := endpoint.NewEnvironmentEndpoint(environmentUseCase)

	return NewEnvironmentMiddleware(environmentEndpoint)
}
