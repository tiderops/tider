package usecase

import (
	"Kubexplorer/internal/model"
	"context"
)

type NodeUseCase interface {
	GetNodes(ctx context.Context, clusterCtx string) ([]model.NodeDto, error)
	GetNode(ctx context.Context, ref model.ResourceRef) (model.NodeDto, error)
}

type nodeUseCase struct {
	client NodeClient
}

func NewNodeUseCase(client NodeClient) NodeUseCase {
	return &nodeUseCase{client: client}
}

func (n *nodeUseCase) GetNodes(ctx context.Context, clusterCtx string) ([]model.NodeDto, error) {
	return n.client.GetNodes(ctx, clusterCtx)
}

func (n *nodeUseCase) GetNode(ctx context.Context, ref model.ResourceRef) (model.NodeDto, error) {
	return n.client.GetNode(ctx, ref)
}
