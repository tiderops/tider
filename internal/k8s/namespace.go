package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"os"
	"path/filepath"
)

type NamespaceClient struct {
	manager ClusterResolver
}

func NewNamespaceClient(manager ClusterResolver) *NamespaceClient {
	return &NamespaceClient{manager: manager}
}

func (n NamespaceClient) GetNamespaces(ctx context.Context, clusterCtx string) ([]model.NamespaceDto, error) {
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	namespaces, err := client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("listing namespaces: %w", err)
	}

	var result []model.NamespaceDto

	for _, ns := range namespaces.Items {
		dto := model.NamespaceDto{
			Name:      ns.Name,
			Age:       model.FormatAge(ns.CreationTimestamp.Time),
			CreatedAt: ns.CreationTimestamp.Unix(),
			Labels:    ns.Labels,
			Status:    string(ns.Status.Phase),
		}
		result = append(result, dto)
	}

	return result, nil
}

func (n NamespaceClient) GetNamespace(ctx context.Context, ref model.ResourceRef) (model.NamespaceDto, error) {
	name, clusterCtx := ref.Name, ref.Cluster
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.NamespaceDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	ns, err := client.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return model.NamespaceDto{}, fmt.Errorf("getting namespace %s: %w", name, err)
	}

	return model.NamespaceDto{
		Name:      ns.Name,
		Age:       model.FormatAge(ns.CreationTimestamp.Time),
		CreatedAt: ns.CreationTimestamp.Unix(),
		Labels:    ns.Labels,
		Status:    string(ns.Status.Phase),
	}, nil
}

func (n NamespaceClient) UpdateNamespace(ctx context.Context, ref model.ResourceRef, dto model.NamespaceDto) error {
	name, clusterCtx := ref.Name, ref.Cluster
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	c := client.CoreV1().Namespaces()
	ns, err := c.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("getting namespace %s: %w", name, err)
	}

	ns.Name = dto.Name

	_, err = c.Update(ctx, ns, metav1.UpdateOptions{})

	return err
}

func (n NamespaceClient) DeleteNamespace(ctx context.Context, ref model.ResourceRef) error {
	name, clusterCtx := ref.Name, ref.Cluster
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	return client.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{})
}

func (n NamespaceClient) ExportObjects(ctx context.Context, clusterCtx string, namespace string, directory string) error {
	client, err := n.manager.ResolveClusterContextDynamic(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	if directory == "" {
		directory = "snapshots"
	}

	targetDir := filepath.Join(directory, namespace)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return fmt.Errorf("creating export directory %s: %w", targetDir, err)
	}

	resourceTypes := []schema.GroupVersionResource{
		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "configmaps"},
		{Group: "", Version: "v1", Resource: "secrets"},
		{Group: "", Version: "v1", Resource: "pods"},
	}

	for _, r := range resourceTypes {
		resList, err := client.Resource(r).Namespace(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("listing %s in namespace %s: %w", r.Resource, namespace, err)
		}

		for _, item := range resList.Items {
			yamlBytes, err := yaml.Marshal(item.Object)
			if err != nil {
				return fmt.Errorf("marshalling %s %s: %w", r.Resource, item.GetName(), err)
			}

			filePath := filepath.Join(targetDir, fmt.Sprintf("%s-%s.yaml", r.Resource, item.GetName()))
			if err := os.WriteFile(filePath, yamlBytes, 0o644); err != nil {
				return fmt.Errorf("writing %s: %w", filePath, err)
			}
		}
	}
	return nil
}
