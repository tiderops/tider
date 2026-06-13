package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"fmt"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobClient struct {
	manager ClusterResolver
}

func NewJob(manager ClusterResolver) *JobClient {
	return &JobClient{
		manager: manager,
	}
}

func (j *JobClient) GetJob(ctx context.Context, ref model.ResourceRef) (*v1.Job, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := j.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("cluster %s is not registered", clusterCtx)
	}

	job, err := client.BatchV1().Jobs(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("getting job %s/%s: %w", namespace, name, err)
	}

	return job, nil
}
