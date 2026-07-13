package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/yaml"
)

type ServiceClient struct {
	manager ClusterResolver
}

func NewServiceClient(manager ClusterResolver) *ServiceClient {
	return &ServiceClient{manager: manager}
}

func (s ServiceClient) GetServices(ctx context.Context, clusterCtx string) ([]model.ServiceDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	services, err := client.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("listing services: %w", err)
	}

	var result []model.ServiceDto

	for _, service := range services.Items {
		var port int32
		if len(service.Spec.Ports) > 0 {
			port = service.Spec.Ports[0].Port
		}

		dto := model.ServiceDto{
			Name:       service.Name,
			Namespace:  service.Namespace,
			Labels:     service.Labels,
			Type:       string(service.Spec.Type),
			InternalIp: service.Spec.ClusterIP,
			ExternalIp: service.Spec.ExternalName,
			Port:       port,
			Status:     service.Status.String(),
			Age:        model.FormatAge(service.CreationTimestamp.Time),
			CreatedAt:  service.CreationTimestamp.Unix(),
			Spec:       service.Spec.String(),
		}

		result = append(result, dto)
	}

	return result, nil
}

func (s ServiceClient) GetService(ctx context.Context, ref model.ResourceRef) (model.ServiceDto, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.ServiceDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	service, err := client.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return model.ServiceDto{}, fmt.Errorf("getting service %s/%s: %w", namespace, name, err)
	}

	return model.ServiceDto{
		Name:      service.Name,
		Namespace: service.Namespace,
		Labels:    service.Labels,
		Status:    service.Status.String(),
		Age:       model.FormatAge(service.CreationTimestamp.Time),
		CreatedAt: service.CreationTimestamp.Unix(),
		Spec:      service.Spec.String(),
	}, nil
}

func (s ServiceClient) UpdateService(ctx context.Context, ref model.ResourceRef, dto model.ServiceUpdate) error {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	c := client.CoreV1().Services(namespace)
	service, err := c.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("getting service %s/%s: %w", namespace, name, err)
	}

	if len(service.Spec.Ports) == 0 {
		return fmt.Errorf("service %s/%s has no ports", namespace, name)
	}

	if service.Labels == nil {
		service.Labels = map[string]string{}
	}
	if service.Spec.Selector == nil {
		service.Spec.Selector = map[string]string{}
	}

	service.Labels["app"] = dto.LabelApp
	service.Spec.Type = v1.ServiceType(dto.SpecType)
	service.Spec.Ports[0].Port = dto.Port
	service.Spec.Ports[0].TargetPort = intstr.FromString(dto.TargetPort)
	service.Spec.Selector["app"] = dto.SelectorApp

	_, err = c.Update(ctx, service, metav1.UpdateOptions{})

	return err
}

func (s ServiceClient) DeleteService(ctx context.Context, ref model.ResourceRef) error {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	return client.CoreV1().Services(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (s ServiceClient) ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := s.manager.ResolveClusterContextDynamic(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	gvr := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "services"}
	res, err := client.Resource(gvr).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return nil, fmt.Errorf("getting service %s/%s: %w", namespace, name, err)
	}

	unstructured.RemoveNestedField(res.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(res.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(res.Object, "metadata", "uid")
	unstructured.RemoveNestedField(res.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(res.Object, "status")

	data, err := yaml.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("marshalling service %s/%s: %w", namespace, name, err)
	}

	return data, nil
}
