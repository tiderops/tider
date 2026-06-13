package usecase

import (
	"Kubexplorer/internal/model"
	"context"
)

type EnvironmentUseCase interface {
	GetCurrentEnvironment(ctx context.Context, env string, name string) (model.EnvironmentDto, error)
	GetObjectsView(ctx context.Context, clusterCtx string) (model.ObjectMapDto, error)
	ListAvailableClusters(ctx context.Context) ([]model.ClusterInfo, error)
}

type environmentUseCase struct {
	client ClusterClient
}

func NewEnvironmentUseCase(client ClusterClient) EnvironmentUseCase {
	return &environmentUseCase{client: client}
}

func (e *environmentUseCase) ListAvailableClusters(ctx context.Context) ([]model.ClusterInfo, error) {
	return e.client.ListAvailableClusters(ctx)
}

func (e *environmentUseCase) GetCurrentEnvironment(ctx context.Context, env string, name string) (model.EnvironmentDto, error) {
	return e.client.GetCurrentCluster(ctx, name)
}

func (e *environmentUseCase) GetObjectsView(ctx context.Context, clusterCtx string) (model.ObjectMapDto, error) {
	objects, err := e.client.GetClusterObjects(ctx, clusterCtx)
	if err != nil {
		return model.ObjectMapDto{}, err
	}

	BuildObjectMap(&objects)

	return objects, nil
}
