package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type nodeClient struct {
	manager ClusterResolver
}

func NewNode(manager ClusterResolver) NodeClient {
	return &nodeClient{manager: manager}
}

func (n *nodeClient) GetNodes(clusterCtx string) ([]model.NodeDtoV2, error) {
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	nodes, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	var result []model.NodeDtoV2

	for _, node := range nodes.Items {
		dto := model.NodeDtoV2{
			Name: node.Name,
			Resource: model.Resource{
				Cpu:     node.Status.Capacity.Cpu().Value(),
				Memory:  node.Status.Capacity.Memory().ScaledValue(resource.Giga),
				Storage: node.Status.Capacity.StorageEphemeral().ScaledValue(resource.Giga),
			},
			KubeletVersion:    node.Status.NodeInfo.KubeletVersion,
			OperatingSystem:   node.Status.NodeInfo.OperatingSystem,
			Version:           node.ResourceVersion,
			CreationTimestamp: node.CreationTimestamp.String(),
			Labels:            node.Labels,
		}

		result = append(result, dto)
	}

	return result, nil
}

func (n *nodeClient) GetNode(name string, clusterCtx string) (model.NodeDtoV2, error) {
	client, err := n.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.NodeDtoV2{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	node, err := client.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	return model.NodeDtoV2{
		Name: node.Name,
		Resource: model.Resource{
			Cpu:     node.Status.Capacity.Cpu().Value(),
			Memory:  node.Status.Capacity.Memory().ScaledValue(resource.Giga),
			Storage: node.Status.Capacity.StorageEphemeral().ScaledValue(resource.Giga),
		},
		Version:           node.ResourceVersion,
		CreationTimestamp: node.CreationTimestamp.String(),
		Labels:            node.Labels,
	}, nil
}
