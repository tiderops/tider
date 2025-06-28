package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type namespaceClient struct {
	manager ClusterResolver
}

func NewNamespaceClient(manager ClusterResolver) NamespaceClient {
	return &namespaceClient{manager: manager}
}

func (n namespaceClient) GetNamespaces(clusterCtx string) ([]model.NamespaceDto, error) {
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	namespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic("Failed to list namespaces")
	}

	var result []model.NamespaceDto

	for _, ns := range namespaces.Items {
		fmt.Println("ns.Name", ns.Name)
		dto := model.NamespaceDto{
			Name:         ns.Name,
			Version:      ns.APIVersion,
			CreationTime: ns.CreationTimestamp.String(),
			Labels:       ns.Labels,
			Status:       ns.Status.String(),
		}
		result = append(result, dto)
	}

	return result, nil
}

func (n namespaceClient) GetNamespace(name string, clusterCtx string) (model.NamespaceDto, error) {
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.NamespaceDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	ns, err := client.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic("Failed to list namespaces")
	}

	return model.NamespaceDto{
		Name:         ns.Name,
		Version:      ns.APIVersion,
		CreationTime: ns.CreationTimestamp.String(),
		Labels:       ns.Labels,
		Status:       ns.Status.String(),
	}, nil
}

func (n namespaceClient) UpdateNamespace(name string, dto model.NamespaceDto, clusterCtx string) error {
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	c := client.CoreV1().Namespaces()
	ns, err := c.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	ns.Name = dto.Name

	_, err = c.Update(context.TODO(), ns, metav1.UpdateOptions{})

	return err
}

func (n namespaceClient) DeleteNamespace(name string, clusterCtx string) error {
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	return client.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
}
