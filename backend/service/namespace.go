package service

import (
	"Kubexplorer/backend/kubeclient"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"os"
)

type NamespaceService interface {
	ExportObjects(namespace string, directory string, clusterCtx string) error
}

type namespaceService struct {
	manager kubeclient.ClusterResolver
}

func NewNamespaceService(manager kubeclient.ClusterResolver) NamespaceService {
	return &namespaceService{manager: manager}
}

func (ns *namespaceService) ExportObjects(namespace string, directory string, clusterCtx string) error {
	client, err := ns.manager.ResolveClusterContextDynamic(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	resourceTypes := []schema.GroupVersionResource{
		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "configmaps"},
		{Group: "", Version: "v1", Resource: "secrets"},
		{Group: "", Version: "v1", Resource: "pods"},
	}

	for _, r := range resourceTypes {
		resList, _ := client.Resource(r).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
		for _, item := range resList.Items {
			yamlBytes, _ := yaml.Marshal(item.Object)
			filePath := fmt.Sprintf("snapshots/%s/%s-%s.yaml", namespace, r.Resource, item.GetName())
			os.WriteFile(filePath, yamlBytes, 0644)
		}
	}
	return nil
}
