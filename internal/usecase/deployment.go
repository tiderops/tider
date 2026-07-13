package usecase

import (
	"Kubexplorer/internal/model"
	"context"
)

type DeploymentUseCase interface {
	GetDeployments(ctx context.Context, clusterCtx string) ([]model.DeploymentDto, error)
	GetDeployment(ctx context.Context, ref model.ResourceRef) (model.DeploymentDto, error)
	UpdateDeployment(ctx context.Context, ref model.ResourceRef, dto model.DeploymentUpdate) error
	DeleteDeployment(ctx context.Context, ref model.ResourceRef) error
	ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error)
}

type deploymentUseCase struct {
	client DeploymentClient
}

func NewDeploymentUseCase(client DeploymentClient) DeploymentUseCase {
	return &deploymentUseCase{client: client}
}

func (d *deploymentUseCase) GetDeployments(ctx context.Context, clusterCtx string) ([]model.DeploymentDto, error) {
	return d.client.GetDeployments(ctx, clusterCtx)
}

func (d *deploymentUseCase) GetDeployment(ctx context.Context, ref model.ResourceRef) (model.DeploymentDto, error) {
	return d.client.GetDeployment(ctx, ref)
}

func (d *deploymentUseCase) UpdateDeployment(ctx context.Context, ref model.ResourceRef, dto model.DeploymentUpdate) error {
	return d.client.UpdateDeployment(ctx, ref, dto)
}

func (d *deploymentUseCase) DeleteDeployment(ctx context.Context, ref model.ResourceRef) error {
	return d.client.DeleteDeployment(ctx, ref)
}

func (d *deploymentUseCase) ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error) {
	return d.client.ExportManifest(ctx, ref)
}
