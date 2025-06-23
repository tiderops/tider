package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math/rand"
	"strconv"
)

type podClient struct {
	manager ClusterResolver
}

func NewPod(manager ClusterResolver) PodClient {
	return &podClient{
		manager: manager,
	}
}

func (p *podClient) GetPods(clusterCtx string) ([]model.PodDto, error) {
	client, err := p.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	podsClient := client.CoreV1().Pods("")

	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error when get pods")
	}
	var podArray []model.PodDto

	for _, pod := range pods.Items {
		age := "0"

		if pod.Status.StartTime != nil {
			age = pod.Status.StartTime.String()
		}

		p := model.PodDto{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Replicas:  1,
			Container: model.Container{
				Limit: model.Resource{
					Cpu:    pod.Spec.Containers[0].Resources.Limits.Cpu().String(),
					Memory: pod.Spec.Containers[0].Resources.Limits.Memory().String(),
				},
				Request: model.Resource{
					Cpu:    pod.Spec.Containers[0].Resources.Requests.Cpu().String(),
					Memory: pod.Spec.Containers[0].Resources.Requests.Memory().String(),
				},
			},
			Status: string(pod.Status.Phase),
			Age:    age,
		}

		podArray = append(podArray, p)

		fmt.Printf("pod name: %s namespace: %s requestCPU: %s limitsCPU: %s requestMemory: %s limitsMemory: %s storage: %s startTime: %s status: %s\n",
			pod.Name,
			pod.Namespace,
			pod.Spec.Containers[0].Resources.Requests.Cpu(),
			pod.Spec.Containers[0].Resources.Limits.Cpu(),
			pod.Spec.Containers[0].Resources.Requests.Memory(),
			pod.Spec.Containers[0].Resources.Limits.Memory(),
			pod.Spec.Containers[0].Resources.Limits.Storage(),
			pod.Status.StartTime,
			pod.Status.Phase,
		)
	}

	fmt.Println("podArray[0].Name", podArray[0].Name)

	return podArray, nil
}

func (p *podClient) GetPod(name string, namespace string, clusterCtx string) (model.PodDto, error) {
	client, err := p.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.PodDto{}, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	pod, _ := client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return model.PodDto{
		Name:      pod.Name,
		Namespace: pod.Namespace,
		Replicas:  1,
		Container: model.Container{
			Limit: model.Resource{
				Cpu:    pod.Spec.Containers[0].Resources.Limits.Cpu().String(),
				Memory: pod.Spec.Containers[0].Resources.Limits.Memory().String(),
			},
			Request: model.Resource{
				Cpu:    pod.Spec.Containers[0].Resources.Requests.Cpu().String(),
				Memory: pod.Spec.Containers[0].Resources.Requests.Memory().String(),
			},
		},
		Status: string(pod.Status.Phase),
		Age:    pod.Status.StartTime.String(),
	}, nil
}

func (p *podClient) UpdatePod(name string, namespace string, dto model.PodDto, clusterCtx string) error {
	client, err := p.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	c := client.CoreV1().Pods(namespace)

	pod, err := c.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	pod.Name = dto.Name
	pod.Namespace = dto.Namespace

	_, err = c.Update(context.TODO(), pod, metav1.UpdateOptions{})

	return err
}

func (p *podClient) DeletePod(name string, namespace string, clusterCtx string) error {
	client, err := p.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	return client.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}
