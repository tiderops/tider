package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"fmt"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

type IngressClient struct {
	manager ClusterResolver
}

func NewIngressClient(manager ClusterResolver) *IngressClient {
	return &IngressClient{manager: manager}
}

func (i IngressClient) GetIngresses(ctx context.Context, clusterCtx string) ([]model.IngressDto, error) {
	client, err := i.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	ingresses, err := client.NetworkingV1().Ingresses("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("listing ingresses: %w", err)
	}

	var result []model.IngressDto

	for _, ingress := range ingresses.Items {
		var rulesDto []model.RuleDto

		for _, rule := range ingress.Spec.Rules {
			dto := model.RuleDto{
				Host:             rule.Host,
				Path:             firstRulePath(rule),
				IngressRuleValue: rule.IngressRuleValue.String(),
			}
			rulesDto = append(rulesDto, dto)
		}

		dto := model.IngressDto{
			Name:      ingress.Name,
			Namespace: ingress.Namespace,
			Age:       model.FormatAge(ingress.CreationTimestamp.Time),
			CreatedAt: ingress.CreationTimestamp.Unix(),
			Labels:    ingress.Labels,
			Rules:     rulesDto,
		}
		result = append(result, dto)
	}

	return result, nil
}

func (i IngressClient) GetIngress(ctx context.Context, ref model.ResourceRef) (model.IngressDto, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := i.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.IngressDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	ingress, err := client.NetworkingV1().Ingresses(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return model.IngressDto{}, fmt.Errorf("getting ingress %s/%s: %w", namespace, name, err)
	}

	var rulesDto []model.RuleDto

	for _, rule := range ingress.Spec.Rules {
		dto := model.RuleDto{
			Host: rule.Host,
			Path: firstRulePath(rule),
		}
		rulesDto = append(rulesDto, dto)
	}

	return model.IngressDto{
		Name:      ingress.GetName(),
		Namespace: ingress.GetNamespace(),
		Age:       model.FormatAge(ingress.GetCreationTimestamp().Time),
		CreatedAt: ingress.GetCreationTimestamp().Unix(),
		Labels:    ingress.GetLabels(),
		Rules:     rulesDto,
	}, nil
}

func (i IngressClient) UpdateIngress(ctx context.Context, ref model.ResourceRef, dto model.IngressDto) error {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := i.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	c := client.NetworkingV1().Ingresses(namespace)
	ingress, err := c.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("getting ingress %s/%s: %w", namespace, name, err)
	}

	ingress.Name = dto.Name
	ingress.Namespace = dto.Namespace
	ingress.Labels = dto.Labels

	_, err = c.Update(ctx, ingress, metav1.UpdateOptions{})

	return err
}

func (i IngressClient) DeleteIngress(ctx context.Context, ref model.ResourceRef) error {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := i.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	return client.NetworkingV1().Ingresses(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (i IngressClient) ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := i.manager.ResolveClusterContextDynamic(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	gvr := schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "ingresses"}
	res, err := client.Resource(gvr).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return nil, fmt.Errorf("getting ingress %s/%s: %w", namespace, name, err)
	}

	unstructured.RemoveNestedField(res.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(res.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(res.Object, "metadata", "uid")
	unstructured.RemoveNestedField(res.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(res.Object, "metadata", "generation")
	unstructured.RemoveNestedField(res.Object, "status")

	data, err := yaml.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("marshalling ingress %s/%s: %w", namespace, name, err)
	}

	return data, nil
}

func firstRulePath(rule networkingv1.IngressRule) string {
	if rule.HTTP == nil || len(rule.HTTP.Paths) == 0 {
		return ""
	}
	return rule.HTTP.Paths[0].Path
}
