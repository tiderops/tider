package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
	"context"
)

type EnvironmentUseCase interface {
	GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error)
	GetObjectsView() (model.ObjectMapDto, error)
	ListAvailableClusters() ([]model.ClusterInfo, error)
}

type environmentUseCase struct {
	client  kubeclient.ClusterClient
	service service.ClusterService
}

func NewEnvironmentUseCase(client kubeclient.ClusterClient) EnvironmentUseCase {
	return &environmentUseCase{client: client}
}

func (e *environmentUseCase) ListAvailableClusters() ([]model.ClusterInfo, error) {
	return e.client.ListAvailableClusters()
}

func (e *environmentUseCase) GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error) {
	return e.client.GetCurrentCluster(name)
}

func (e *environmentUseCase) GetObjectsView() (model.ObjectMapDto, error) {
	return e.service.BuildPodRelationships(context.Background())
}
