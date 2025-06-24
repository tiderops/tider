package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type serviceClient struct {
	manager ClusterResolver
}

func NewServiceClient(manager ClusterResolver) ServiceClient {
	return &serviceClient{manager: manager}
}

func (s serviceClient) GetServices(clusterCtx string) ([]model.ServiceDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	services, err := client.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic("Failed to list services")
	}

	var result []model.ServiceDto

	for _, service := range services.Items {
		dto := model.ServiceDto{
			Name:              service.Name,
			Namespace:         service.Namespace,
			Labels:            service.Labels,
			Status:            service.Status.String(),
			CreationTimestamp: service.CreationTimestamp.String(),
			Spec:              service.Spec.String(),
		}

		result = append(result, dto)
	}

	return result, errors.New("No services found")
}

func (s serviceClient) GetService(name string, namespace string, clusterCtx string) (model.ServiceDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.ServiceDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	service, err := client.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic("Failed to get service")
	}

	return model.ServiceDto{
		Name:              service.Name,
		Namespace:         service.Namespace,
		Labels:            service.Labels,
		Status:            service.Status.String(),
		CreationTimestamp: service.CreationTimestamp.String(),
		Spec:              service.Spec.String(),
	}, errors.New("No service found")
}

func (s serviceClient) UpdateService(name string, namespace string, dto model.ServiceDto, clusterCtx string) error {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	c := client.CoreV1().Services(namespace)
	service, err := c.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	service.Name = dto.Name
	service.Namespace = dto.Namespace
	service.Labels = dto.Labels

	_, err = c.Update(context.TODO(), service, metav1.UpdateOptions{})

	return err
}

func (s serviceClient) DeleteService(name string, namespace string, clusterCtx string) error {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	return client.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}
