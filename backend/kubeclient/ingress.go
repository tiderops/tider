package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

type ingressClient struct {
	manager ClusterResolver
}

func NewIngressClient(manager ClusterResolver) IngressClient {
	return &ingressClient{manager: manager}
}

func (i ingressClient) GetIngresses(clusterCtx string) ([]model.IngressDto, error) {
	client, err := i.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	ingresses, err := client.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic("Error while listing ingresses")
	}

	var result []model.IngressDto

	for _, ingress := range ingresses.Items {
		var rulesDto []model.RuleDto

		for _, rule := range ingress.Spec.Rules {
			dto := model.RuleDto{
				Host:             rule.Host,
				Path:             rule.HTTP.Paths[0].Path,
				IngressRuleValue: rule.IngressRuleValue.String(),
			}
			rulesDto = append(rulesDto, dto)
		}

		dto := model.IngressDto{
			Name:      ingress.Name,
			Namespace: ingress.Namespace,
			Age:       ingress.CreationTimestamp.String(),
			Labels:    ingress.Labels,
			Rules:     rulesDto,
		}
		result = append(result, dto)
	}

	return result, nil
}

func (i ingressClient) GetIngress(name string, namespace string, clusterCtx string) (model.IngressDto, error) {
	client, err := i.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.IngressDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	ingress, _ := client.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	var rulesDto []model.RuleDto

	for _, rule := range ingress.Spec.Rules {
		dto := model.RuleDto{
			Host: rule.Host,
			Path: rule.HTTP.Paths[0].Path,
		}
		rulesDto = append(rulesDto, dto)
	}

	return model.IngressDto{
		Name:      ingress.GetName(),
		Namespace: ingress.GetNamespace(),
		Age:       ingress.GetCreationTimestamp().String(),
		Labels:    ingress.GetLabels(),
		Rules:     rulesDto,
	}, nil
}

func (i ingressClient) UpdateIngress(name string, namespace string, dto model.IngressDto, clusterCtx string) error {
	client, err := i.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	c := client.NetworkingV1().Ingresses(namespace)
	ingress, err := c.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	ingress.Name = dto.Name
	ingress.Namespace = dto.Namespace
	ingress.Labels = dto.Labels

	_, err = c.Update(context.TODO(), ingress, metav1.UpdateOptions{})

	return err
}

func (i ingressClient) DeleteIngress(name string, namespace string, clusterCtx string) error {
	client, err := i.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	return client.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (i ingressClient) ExportManifest(name string, namespace string, clusterCtx string) ([]byte, error) {
	client, err := i.manager.ResolveClusterContextDynamic(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	gvr := schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "ingresses"}
	res, err := client.Resource(gvr).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		return nil, fmt.Errorf("Error when get ingresses %s in namespace %s\n", name, namespace)
	}

	unstructured.RemoveNestedField(res.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(res.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(res.Object, "metadata", "uid")
	unstructured.RemoveNestedField(res.Object, "metadata", "selfLink")
	unstructured.RemoveNestedField(res.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(res.Object, "metadata", "generation")
	unstructured.RemoveNestedField(res.Object, "status")

	data, _ := yaml.Marshal(res)
	fmt.Println(string(data))
	return data, nil
}
