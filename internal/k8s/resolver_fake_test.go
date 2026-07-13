package k8s

import (
	"errors"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

// fakeResolver serves one fake clientset for any cluster name
type fakeResolver struct {
	client kubernetes.Interface
}

func (f *fakeResolver) ResolveClusterContext(name string) (kubernetes.Interface, error) {
	if f.client == nil {
		return nil, errors.New("cluster is not registered")
	}
	return f.client, nil
}

func (f *fakeResolver) ResolveClusterContextDynamic(name string) (dynamic.Interface, error) {
	return nil, errors.New("dynamic client not faked")
}

func (f *fakeResolver) ResolveClusterMetric(name string) (metricsv.Interface, error) {
	return nil, errors.New("metric client not faked")
}
