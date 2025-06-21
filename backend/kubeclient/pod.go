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

type podClient struct {
	client kubernetes.Interface
}

func NewPod(client kubernetes.Interface) PodClient {
	return &podClient{client: client}
}

func (p *podClient) GetPodsMock() ([]model.PodDto, error) {
	var pods []model.PodDto
	for i := 0; i < 10; i++ {

		p := model.PodDto{
			Name:      fmt.Sprintf("pod %d", i),
			Namespace: "TODO",
			Replicas:  1,
			Container: model.Container{
				Limit: model.Resource{
					Cpu:    strconv.Itoa(rand.Intn(1000)) + "mi",
					Memory: strconv.Itoa(rand.Intn(1000)),
				},
				Request: model.Resource{
					Cpu:    strconv.Itoa(rand.Intn(1000)),
					Memory: strconv.Itoa(rand.Intn(1000)),
				},
			},
			Status: "Alive",
			Age:    strconv.Itoa(rand.Intn(1000)),
		}

		pods = append(pods, p)
	}

	fmt.Println(pods[0].Name)

	return pods, errors.New("Not implemented")
}

func (p *podClient) GetPods(clusterCtx string) ([]model.PodDto, error) {
	//podsClient := p.client.CoreV1().Pods("")
	client, err := GlobalClusterManager.GetValue(clusterCtx)
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

func (p *podClient) GetPod(name string, namespace string) (model.PodDto, error) {
	pod, _ := p.client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})

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
	}, errors.New("Error while listing ingresses")
}

func (p *podClient) UpdatePod(name string, namespace string, dto model.PodDto) error {
	client := p.client.CoreV1().Pods(namespace)
	pod, err := client.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	pod.Name = dto.Name
	pod.Namespace = dto.Namespace

	_, err = client.Update(context.TODO(), pod, metav1.UpdateOptions{})

	return err
}

func (p *podClient) DeletePod(name string, namespace string) error {
	return p.client.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

//func status(namespace string) {
//	podsClient := client.CoreV1().Pods(namespace)
//	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		fmt.Println("Error to get pods")
//	}
//
//	for _, pod := range pods.Items {
//
//		if pod.Status.ContainerStatuses[0].State.Waiting != nil && pod.Status.ContainerStatuses[0].State.Waiting.Reason == string(knowledge.CRASH_LOOP_BACK_OFF) {
//			fmt.Printf("Pod %s is in CrashLoopBackOff: %s\n", pod.Name, pod.Status.ContainerStatuses[0].State.Waiting.Reason)
//		}
//		fmt.Printf("pod: %s\n", pod.Status.ContainerStatuses[0].State.Waiting)
//	}
//
//	knowledge.ErrorSource(knowledge.PODS, string(knowledge.CRASH_LOOP_BACK_OFF))
//}

// ExampleMapper
//func mapPod(n v1.Node) model.PodDto {
//	return model.PodDto{
//		Name: n.Name,
//	}
//}
