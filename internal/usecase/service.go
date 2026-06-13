package usecase

import (
	"Kubexplorer/internal/model"
	"context"
)

type ServiceUseCase interface {
	GetServices(ctx context.Context, clusterCtx string) ([]model.ServiceDto, error)
	GetService(ctx context.Context, ref model.ResourceRef) (model.ServiceDto, error)
	UpdateService(ctx context.Context, ref model.ResourceRef, dto model.ServiceUpdate) error
	DeleteService(ctx context.Context, ref model.ResourceRef) error
	ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error)
}

type serviceUseCase struct {
	client ServiceClient
}

func NewServiceUseCase(client ServiceClient) ServiceUseCase {
	return &serviceUseCase{client: client}
}

func (s *serviceUseCase) GetServices(ctx context.Context, clusterCtx string) ([]model.ServiceDto, error) {
	return s.client.GetServices(ctx, clusterCtx)
}

func (s *serviceUseCase) GetService(ctx context.Context, ref model.ResourceRef) (model.ServiceDto, error) {
	return s.client.GetService(ctx, ref)
}

func (s *serviceUseCase) UpdateService(ctx context.Context, ref model.ResourceRef, dto model.ServiceUpdate) error {
	return s.client.UpdateService(ctx, ref, dto)
}

func (s *serviceUseCase) DeleteService(ctx context.Context, ref model.ResourceRef) error {
	return s.client.DeleteService(ctx, ref)
}

func (s *serviceUseCase) ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error) {
	return s.client.ExportManifest(ctx, ref)
}
