package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type NodeUseCase interface {
	GetNodes() ([]model.NodeDtoV2, error)
	GetNode(name string) (model.NodeDtoV2, error)
}

type nodeUseCase struct {
	client kubeclient.ClusterClient
}

func NewNodeUseCase(client kubeclient.ClusterClient) NodeUseCase {
	return &nodeUseCase{client: client}
}

func (n *nodeUseCase) GetNodes() ([]model.NodeDtoV2, error) {
	return n.client.GetNodes()
}

func (n *nodeUseCase) GetNode(name string) (model.NodeDtoV2, error) {
	return n.client.GetNode(name)
}
