package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type GeneralEndpoint struct {
	nodeUseCase      usecase.NodeUseCase
	namespaceUseCase usecase.NamespaceUseCase
}

func NewGeneralEndpoint(nodeUseCase usecase.NodeUseCase, namespaceUseCase usecase.NamespaceUseCase) *GeneralEndpoint {
	return &GeneralEndpoint{
		nodeUseCase:      nodeUseCase,
		namespaceUseCase: namespaceUseCase,
	}
}

func (ge *GeneralEndpoint) GetNodes(clusterCtx string) ([]model.NodeDtoV2, error) {
	return ge.nodeUseCase.GetNodes(clusterCtx)
}

func (ge *GeneralEndpoint) GetNode(name string, clusterCtx string) (model.NodeDtoV2, error) {
	return ge.nodeUseCase.GetNode(name, clusterCtx)
}

func (ge *GeneralEndpoint) GetNamespaces(clusterCtx string) ([]model.NamespaceDto, error) {
	return ge.namespaceUseCase.GetNamespaces(clusterCtx)
}

func (ge *GeneralEndpoint) GetNamespace(name string, clusterCtx string) (model.NamespaceDto, error) {
	return ge.namespaceUseCase.GetNamespace(name, clusterCtx)
}

func (ge *GeneralEndpoint) UpdateNamespace(name string, dto model.NamespaceDto, clusterCtx string) error {
	return ge.namespaceUseCase.UpdateNamespace(name, dto, clusterCtx)
}

func (ge *GeneralEndpoint) DeleteNamespace(name string, clusterCtx string) error {
	return ge.namespaceUseCase.DeleteNamespace(name, clusterCtx)
}

func (ge *GeneralEndpoint) ExportNamespaceObjects(namespace string, directory string, clusterCtx string) error {
	return ge.namespaceUseCase.ExportNamespaceObjects(namespace, directory, clusterCtx)
}
