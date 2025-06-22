package middleware

import (
	"Kubexplorer/backend/endpoint"
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type GeneralMiddleware struct {
	endpoint endpoint.GeneralEndpoint
}

func NewGeneralMiddleware(endpoint *endpoint.GeneralEndpoint) *GeneralMiddleware {
	return &GeneralMiddleware{endpoint: *endpoint}
}

func (g *GeneralMiddleware) GetNodes() ([]model.NodeDtoV2, error) {
	return g.endpoint.GetNodes()
}

func (g *GeneralMiddleware) GetNode(name string) (model.NodeDtoV2, error) {
	return g.endpoint.GetNode(name)
}

func (g *GeneralMiddleware) GetNamespaces() ([]model.NamespaceDto, error) {
	return g.endpoint.GetNamespaces()
}

func (g *GeneralMiddleware) GetNamespace(name string) (model.NamespaceDto, error) {
	return g.endpoint.GetNamespace(name)
}

func (g *GeneralMiddleware) UpdateNamespace(name string, dto model.NamespaceDto) error {
	return g.endpoint.UpdateNamespace(name, dto)
}

func (g *GeneralMiddleware) DeleteNamespace(name string) error {
	return g.endpoint.DeleteNamespace(name)
}

func (g *GeneralMiddleware) ExportNamespaceObjects(namespace string, directory string) error {
	return g.endpoint.ExportNamespaceObjects(namespace, directory)
}

func BuildGeneral() *GeneralMiddleware {
	//x, _ := kubeclient.GlobalClusterManager.GetValue("as")

	nodeClient := kubeclient.NewCluster()
	namespaceClient := kubeclient.NewNamespaceClient(nil)

	nodeUseCase := usecase.NewNodeUseCase(nodeClient)
	namespaceUseCase := usecase.NewNamespaceUseCase(namespaceClient)

	generalEndpoint := endpoint.NewGeneralEndpoint(nodeUseCase, namespaceUseCase)

	return NewGeneralMiddleware(generalEndpoint)
}
