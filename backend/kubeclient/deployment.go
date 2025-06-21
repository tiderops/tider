package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"math/rand"
	"strconv"
)

type deploymentClient struct {
	client kubernetes.Interface
}

func NewDeployment(client kubernetes.Interface) DeploymentClient {
	return &deploymentClient{client: client}
}

func (d deploymentClient) GetDeployments(clusterCtx string) ([]model.DeploymentDto, error) {
	client, err := GlobalClusterManager.GetValue(clusterCtx)
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

	return result, errors.New("deployment list is empty")
}

func (d deploymentClient) GetDeploymentsMock() ([]model.DeploymentDto, error) {
	var deployments []model.DeploymentDto
	for i := 0; i < 10; i++ {

		p := model.DeploymentDto{
			Name:      fmt.Sprintf("Deployment %d", i),
			Namespace: "TODO",
			Status:    "Running",
			Age:       strconv.Itoa(rand.Intn(1000)),
		}

		deployments = append(deployments, p)
	}

	return deployments, errors.New("Deployment Not Found")
}

func (d deploymentClient) GetDeployment(name string, namespace string) (model.DeploymentDto, error) {
	deployment, _ := d.client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return model.DeploymentDto{
		Name:      deployment.Name,
		Namespace: deployment.Namespace,
		Status:    string(deployment.Status.Conditions[0].Status),
		Age:       deployment.CreationTimestamp.String(),
	}, errors.New("Error while listing ingresses")
}

func (d deploymentClient) UpdateDeployment(name string, namespace string, dto model.DeploymentDto) error {
	client := d.client.AppsV1().Deployments(namespace)
	deployment, err := client.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	deployment.Name = dto.Name
	deployment.Namespace = dto.Namespace

	_, err = client.Update(context.TODO(), deployment, metav1.UpdateOptions{})

	return err
}

func (d deploymentClient) DeleteDeployment(name string, namespace string) error {
	return d.client.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}
