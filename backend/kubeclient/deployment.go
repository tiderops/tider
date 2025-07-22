package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

type deploymentClient struct {
	manager ClusterResolver
}

func NewDeployment(manager ClusterResolver) DeploymentClient {
	return &deploymentClient{manager: manager}
}

func (d deploymentClient) GetDeployments(clusterCtx string) ([]model.DeploymentDto, error) {
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	deployments, err := client.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic("Error when get deploys")
	}
	var result []model.DeploymentDto

	for _, deploy := range deployments.Items {
		d := model.DeploymentDto{
			Name:      deploy.Name,
			Namespace: deploy.Namespace,
			Status:    string(deploy.Status.Conditions[0].Status),
			Age:       deploy.CreationTimestamp.String(),
		}

		result = append(result, d)

		fmt.Printf("dep name: %s namespace: %s status: %s startTime: %s\n",
			deploy.Name,
			deploy.Namespace,
			deploy.Status.Conditions[0].Status,
			deploy.CreationTimestamp,
		)
	}

	return result, nil
}

func (d deploymentClient) GetDeployment(name string, namespace string, clusterCtx string) (model.DeploymentDto, error) {
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.DeploymentDto{}, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	deployment, _ := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return model.DeploymentDto{
		Name:      deployment.Name,
		Namespace: deployment.Namespace,
		Status:    string(deployment.Status.Conditions[0].Status),
		Age:       deployment.CreationTimestamp.String(),
	}, nil
}

func (d deploymentClient) UpdateDeployment(name string, namespace string, dto model.DeploymentRequest, clusterCtx string) error {
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	c := client.AppsV1().Deployments(namespace)
	deployment, err := c.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	replicas, _ := strconv.Atoi(dto.Replicas)
	ptrReplica := int32(replicas)
	strategyType := v1.DeploymentStrategyType(dto.StrategyType)

	deployment.Spec.Replicas = &ptrReplica
	deployment.Spec.Selector.MatchLabels["app"] = dto.App
	deployment.Spec.Strategy.Type = strategyType
	deployment.Spec.Template.Labels["app"] = dto.Label.App
	deployment.Spec.Template.Labels["tier"] = dto.Label.Tier
	deployment.Spec.Template.Labels["tierType"] = dto.Label.TierType

	_, err = c.Update(context.TODO(), deployment, metav1.UpdateOptions{})

	return err
}

func (d deploymentClient) DeleteDeployment(name string, namespace string, clusterCtx string) error {
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	return client.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}
