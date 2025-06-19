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

func (ge *GeneralEndpoint) GetNodes() ([]model.NodeDtoV2, error) {
	return ge.nodeUseCase.GetNodes()
}

func (ge *GeneralEndpoint) GetNode(name string) (model.NodeDtoV2, error) {
	return ge.nodeUseCase.GetNode(name)
}

func (ge *GeneralEndpoint) GetNamespaces() ([]model.NamespaceDto, error) {
	return ge.namespaceUseCase.GetNamespaces()
}

func (ge *GeneralEndpoint) GetNamespace(name string) (model.NamespaceDto, error) {
	return ge.namespaceUseCase.GetNamespace(name)
}

func (ge *GeneralEndpoint) UpdateNamespace(name string, dto model.NamespaceDto) error {
	return ge.namespaceUseCase.UpdateNamespace(name, dto)
}

func (ge *GeneralEndpoint) DeleteNamespace(name string) error {
	return ge.namespaceUseCase.DeleteNamespace(name)
}

func (ge *GeneralEndpoint) ExportNamespaceObjects(namespace string, directory string) error {
	return ge.namespaceUseCase.ExportNamespaceObjects(namespace, directory)
}
