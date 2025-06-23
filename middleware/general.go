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

func (g *GeneralMiddleware) GetNodes(clusterCtx string) ([]model.NodeDtoV2, error) {
	return g.endpoint.GetNodes(clusterCtx)
}

func (g *GeneralMiddleware) GetNode(name string, clusterCtx string) (model.NodeDtoV2, error) {
	return g.endpoint.GetNode(name, clusterCtx)
}

func (g *GeneralMiddleware) GetNamespaces(clusterCtx string) ([]model.NamespaceDto, error) {
	return g.endpoint.GetNamespaces(clusterCtx)
}

func (g *GeneralMiddleware) GetNamespace(name string, clusterCtx string) (model.NamespaceDto, error) {
	return g.endpoint.GetNamespace(name, clusterCtx)
}

func (g *GeneralMiddleware) UpdateNamespace(name string, dto model.NamespaceDto, clusterCtx string) error {
	return g.endpoint.UpdateNamespace(name, dto, clusterCtx)
}

func (g *GeneralMiddleware) DeleteNamespace(name string, clusterCtx string) error {
	return g.endpoint.DeleteNamespace(name, clusterCtx)
}

func (g *GeneralMiddleware) ExportNamespaceObjects(namespace string, directory string, clusterCtx string) error {
	return g.endpoint.ExportNamespaceObjects(namespace, directory, clusterCtx)
}

func BuildGeneral(manager kubeclient.ClusterResolver) *GeneralMiddleware {
	nodeClient := kubeclient.NewNode(manager)
	namespaceClient := kubeclient.NewNamespaceClient(manager)

	nodeUseCase := usecase.NewNodeUseCase(nodeClient)
	namespaceUseCase := usecase.NewNamespaceUseCase(namespaceClient)

	generalEndpoint := endpoint.NewGeneralEndpoint(nodeUseCase, namespaceUseCase)

	return NewGeneralMiddleware(generalEndpoint)
}
