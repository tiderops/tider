package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"fmt"
	v1_apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

type DeploymentClient struct {
	manager ClusterResolver
}

func NewDeployment(manager ClusterResolver) *DeploymentClient {
	return &DeploymentClient{manager: manager}
}

func (d DeploymentClient) GetDeployments(ctx context.Context, clusterCtx string) ([]model.DeploymentDto, error) {
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	deployments, err := client.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("listing deployments: %w", err)
	}

	var result []model.DeploymentDto

	for _, deploy := range deployments.Items {
		dto := model.DeploymentDto{
			Name:      deploy.Name,
			Namespace: deploy.Namespace,
			Replicas:  deploy.Status.Replicas,
			Age:       model.FormatAge(deploy.CreationTimestamp.Time),
			CreatedAt: deploy.CreationTimestamp.Unix(),
			Status:    deploymentStatus(deploy),
			Labels:    deploy.Labels,
		}

		result = append(result, dto)
	}

	return result, nil
}

func (d DeploymentClient) ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := d.manager.ResolveClusterContextDynamic(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	gvr := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	res, err := client.Resource(gvr).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("getting deployment %s/%s: %w", namespace, name, err)
	}

	unstructured.RemoveNestedField(res.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(res.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(res.Object, "metadata", "uid")
	unstructured.RemoveNestedField(res.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(res.Object, "metadata", "generation")
	unstructured.RemoveNestedField(res.Object, "status")

	data, err := yaml.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("marshalling deployment %s/%s: %w", namespace, name, err)
	}

	return data, nil
}

func (d DeploymentClient) GetDeployment(ctx context.Context, ref model.ResourceRef) (model.DeploymentDto, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.DeploymentDto{}, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return model.DeploymentDto{}, fmt.Errorf("getting deployment %s/%s: %w", namespace, name, err)
	}

	return model.DeploymentDto{
		Name:       deployment.Name,
		Namespace:  deployment.Namespace,
		Containers: containersInfo(deployment.Spec.Template.Spec.Containers),
		Status:     deploymentStatus(*deployment),
		Age:        model.FormatAge(deployment.CreationTimestamp.Time),
		CreatedAt:  deployment.CreationTimestamp.Unix(),
	}, nil
}

func (d DeploymentClient) GetDeploymentObject(ctx context.Context, ref model.ResourceRef) (*v1_apps.Deployment, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("getting deployment %s/%s: %w", namespace, name, err)
	}

	return deployment, nil
}

func (d DeploymentClient) UpdateDeployment(ctx context.Context, ref model.ResourceRef, dto model.DeploymentUpdate) error {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	c := client.AppsV1().Deployments(namespace)
	deployment, err := c.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("getting deployment %s/%s: %w", namespace, name, err)
	}

	if len(deployment.Spec.Template.Spec.Containers) == 0 {
		return fmt.Errorf("deployment %s/%s has no containers", namespace, name)
	}

	// TODO: finish update feature (replicas, strategy, labels)
	cpu := resource.MustParse(fmt.Sprintf("%dm", dto.Container.Resource.LCpu))
	mem := resource.MustParse(fmt.Sprintf("%dMi", dto.Container.Resource.LMemory))

	limits := deployment.Spec.Template.Spec.Containers[0].Resources.Limits
	if limits == nil {
		limits = v1.ResourceList{}
		deployment.Spec.Template.Spec.Containers[0].Resources.Limits = limits
	}
	limits[v1.ResourceCPU] = cpu
	limits[v1.ResourceMemory] = mem

	_, err = c.Update(ctx, deployment, metav1.UpdateOptions{})

	return err
}

func (d DeploymentClient) DeleteDeployment(ctx context.Context, ref model.ResourceRef) error {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := d.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	return client.AppsV1().Deployments(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func deploymentStatus(deploy v1_apps.Deployment) string {
	if len(deploy.Status.Conditions) == 0 {
		return "Unknown"
	}
	return string(deploy.Status.Conditions[0].Status)
}
