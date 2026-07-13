package usecase

import (
	"Kubexplorer/internal/model"
	"context"
)

type NamespaceUseCase interface {
	GetNamespaces(ctx context.Context, clusterCtx string) ([]model.NamespaceDto, error)
	GetNamespace(ctx context.Context, ref model.ResourceRef) (model.NamespaceDto, error)
	UpdateNamespace(ctx context.Context, ref model.ResourceRef, dto model.NamespaceDto) error
	DeleteNamespace(ctx context.Context, ref model.ResourceRef) error
	ExportNamespaceObjects(ctx context.Context, clusterCtx string, namespace string, directory string) error
}

type namespaceUseCase struct {
	client NamespaceClient
}

func NewNamespaceUseCase(client NamespaceClient) NamespaceUseCase {
	return &namespaceUseCase{client: client}
}

func (n *namespaceUseCase) GetNamespaces(ctx context.Context, clusterCtx string) ([]model.NamespaceDto, error) {
	return n.client.GetNamespaces(ctx, clusterCtx)
}
func (n *namespaceUseCase) GetNamespace(ctx context.Context, ref model.ResourceRef) (model.NamespaceDto, error) {
	return n.client.GetNamespace(ctx, ref)
}

func (n *namespaceUseCase) UpdateNamespace(ctx context.Context, ref model.ResourceRef, dto model.NamespaceDto) error {
	return n.client.UpdateNamespace(ctx, ref, dto)
}

func (n *namespaceUseCase) DeleteNamespace(ctx context.Context, ref model.ResourceRef) error {
	return n.client.DeleteNamespace(ctx, ref)
}

func (n *namespaceUseCase) ExportNamespaceObjects(ctx context.Context, clusterCtx string, namespace string, directory string) error {
	return n.client.ExportObjects(ctx, clusterCtx, namespace, directory)
}
