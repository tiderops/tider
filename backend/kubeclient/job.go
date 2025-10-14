package kubeclient

import (
	"context"
	"fmt"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type jobClient struct {
	manager ClusterResolver
}

func NewJob(manager ClusterResolver) JobClient {
	return &jobClient{
		manager: manager,
	}
}

func (j *jobClient) GetJob(name string, namespace string, clusterCtx string) (*v1.Job, error) {
	client, err := j.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return &v1.Job{}, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	job, _ := client.BatchV1().Jobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return job, nil
}
