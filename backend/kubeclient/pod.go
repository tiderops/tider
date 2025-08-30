package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
			Container: model.Container{
				//Image:      pod.Spec.Containers[0].Image,
				PullPolicy: string(pod.Spec.Containers[0].ImagePullPolicy),
				//Port:       string(pod.Spec.Containers[0].Ports[0].ContainerPort),
				Limit: model.Resource{
					Cpu:    pod.Spec.Containers[0].Resources.Limits.Cpu().MilliValue(),
					Memory: pod.Spec.Containers[0].Resources.Limits.Memory().ScaledValue(resource.Mega),
				},
				Request: model.Resource{
					Cpu:    pod.Spec.Containers[0].Resources.Requests.Cpu().MilliValue(),
					Memory: pod.Spec.Containers[0].Resources.Requests.Memory().ScaledValue(resource.Mega),
				},
			},
			Node:   pod.Spec.NodeName,
			Status: string(pod.Status.Phase),
			Age:    age,
			Labels: pod.Labels,
		}

		podArray = append(podArray, p)
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

	age := "0"
	var port int32 = 0

	if pod.Status.StartTime != nil {
		age = pod.Status.StartTime.String()
	}

	for _, c := range pod.Spec.Containers {
		for _, pr := range c.Ports {
			fmt.Printf("Port: %d\n", pr.ContainerPort)
			port = pr.ContainerPort
		}
	}

	dto := model.PodDto{
		Name:      pod.Name,
		Namespace: pod.Namespace,
		Container: model.Container{
			Image:      pod.Spec.Containers[0].Image,
			PullPolicy: string(pod.Spec.Containers[0].ImagePullPolicy),
			Port:       port,
			Limit: model.Resource{
				Cpu:    pod.Spec.Containers[0].Resources.Limits.Cpu().MilliValue(),
				Memory: pod.Spec.Containers[0].Resources.Limits.Memory().ScaledValue(resource.Mega),
			},
			Request: model.Resource{
				Cpu:    pod.Spec.Containers[0].Resources.Requests.Cpu().MilliValue(),
				Memory: pod.Spec.Containers[0].Resources.Requests.Memory().ScaledValue(resource.Mega),
			},
		},
		Status:   string(pod.Status.Phase),
		Age:      age,
		Editable: []string{"name", "namespace", "image", "pullPolicy", "port", "RMemory", "RCpu", "LMemory", "LCpu"},
	}

	return dto, nil
}

func (p *podClient) UpdatePod(name string, namespace string, dto model.PodUpdate, clusterCtx string) error {
	client, err := p.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	c := client.CoreV1().Pods(namespace)

	pod, err := c.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching pod")
	}

	//port, _ := strconv.ParseInt(dto.Container.Port, 10, 32)
	//rCpu, _ := strconv.ParseInt(dto.Container.Resource.RCpu, 10, 32)
	//rMemory, _ := strconv.ParseInt(dto.Container.Resource.RMemory, 10, 32)
	//lCpu, _ := strconv.ParseInt(dto.Container.Resource.LCpu, 10, 32)
	//lMemory, _ := strconv.ParseInt(dto.Container.Resource.LMemory, 10, 32)
	//pullPolicy := v1.PullPolicy(dto.Container.PullPolicy)

	pod.ObjectMeta.Labels["app"] = dto.App
	pod.Spec.Containers[0].Image = dto.Container.Image
	//pod.Spec.Containers[0].ImagePullPolicy = pullPolicy
	//pod.Spec.Containers[0].Ports[0].ContainerPort = int32(port)
	//pod.Spec.Containers[0].Resources.Requests.Cpu().Set(rCpu)
	//pod.Spec.Containers[0].Resources.Requests.Memory().Set(rMemory)
	//pod.Spec.Containers[0].Resources.Limits.Cpu().Set(lCpu)
	//pod.Spec.Containers[0].Resources.Limits.Memory().Set(lMemory)

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
