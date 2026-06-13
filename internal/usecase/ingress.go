package usecase

import (
	"Kubexplorer/internal/model"
	"context"
)

type IngressUseCase interface {
	GetIngresses(ctx context.Context, clusterCtx string) ([]model.IngressDto, error)
	GetIngress(ctx context.Context, ref model.ResourceRef) (model.IngressDto, error)
	UpdateIngress(ctx context.Context, ref model.ResourceRef, dto model.IngressDto) error
	DeleteIngress(ctx context.Context, ref model.ResourceRef) error
	ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error)
}

type ingressUseCase struct {
	client IngressClient
}

func NewIngressUseCase(client IngressClient) IngressUseCase {
	return &ingressUseCase{client: client}
}

func (i *ingressUseCase) GetIngresses(ctx context.Context, clusterCtx string) ([]model.IngressDto, error) {
	return i.client.GetIngresses(ctx, clusterCtx)
}

func (i *ingressUseCase) GetIngress(ctx context.Context, ref model.ResourceRef) (model.IngressDto, error) {
	return i.client.GetIngress(ctx, ref)
}

func (i *ingressUseCase) UpdateIngress(ctx context.Context, ref model.ResourceRef, dto model.IngressDto) error {
	return i.client.UpdateIngress(ctx, ref, dto)
}

func (i *ingressUseCase) DeleteIngress(ctx context.Context, ref model.ResourceRef) error {
	return i.client.DeleteIngress(ctx, ref)
}

func (i *ingressUseCase) ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error) {
	return i.client.ExportManifest(ctx, ref)
}
