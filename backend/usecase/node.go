package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type NodeUseCase interface {
	GetNodes(clusterCtx string) ([]model.NodeDtoV2, error)
	GetNode(name string, clusterCtx string) (model.NodeDtoV2, error)
}

type nodeUseCase struct {
	client kubeclient.NodeClient
}

func NewNodeUseCase(client kubeclient.NodeClient) NodeUseCase {
	return &nodeUseCase{client: client}
}

func (n *nodeUseCase) GetNodes(clusterCtx string) ([]model.NodeDtoV2, error) {
	return n.client.GetNodes(clusterCtx)
}

func (n *nodeUseCase) GetNode(name string, clusterCtx string) (model.NodeDtoV2, error) {
	return n.client.GetNode(name, clusterCtx)
}
