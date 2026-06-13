package k8s

import (
	"Kubexplorer/internal/apperr"
	"fmt"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
	"sync"
)

type ClusterResolver interface {
	ResolveClusterContext(name string) (kubernetes.Interface, error)
	ResolveClusterContextDynamic(name string) (dynamic.Interface, error)
	ResolveClusterMetric(name string) (metricsv.Interface, error)
}

type ClusterManager struct {
	mu             sync.RWMutex
	clients        map[string]*kubernetes.Clientset
	dynamicClients map[string]*dynamic.DynamicClient
	metricClients  map[string]*metricsv.Clientset
}

func NewClusterManager() *ClusterManager {
	return &ClusterManager{
		clients:        make(map[string]*kubernetes.Clientset),
		dynamicClients: make(map[string]*dynamic.DynamicClient),
		metricClients:  make(map[string]*metricsv.Clientset),
	}
}

func (cm *ClusterManager) ResolveClusterContext(name string) (kubernetes.Interface, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	if client, ok := cm.clients[name]; ok {
		return client, nil
	}
	return nil, apperr.New(apperr.KindUnreachable, fmt.Sprintf("cluster %q is not registered", name))
}

func (cm *ClusterManager) ResolveClusterContextDynamic(name string) (dynamic.Interface, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	if client, ok := cm.dynamicClients[name]; ok {
		return client, nil
	}
	return nil, apperr.New(apperr.KindUnreachable, fmt.Sprintf("cluster %q is not registered", name))
}

func (cm *ClusterManager) ResolveClusterMetric(name string) (metricsv.Interface, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	if client, ok := cm.metricClients[name]; ok {
		return client, nil
	}
	return nil, apperr.New(apperr.KindUnreachable, fmt.Sprintf("cluster %q is not registered", name))
}

func (cm *ClusterManager) GetClient(clusterName string, kubeConfigPath string) (*kubernetes.Clientset, error) {
	cm.mu.RLock()
	if client, exists := cm.clients[clusterName]; exists {
		cm.mu.RUnlock()
		return client, nil
	}
	cm.mu.RUnlock()

	cm.mu.Lock()
	defer cm.mu.Unlock()

	if client, exists := cm.clients[clusterName]; exists {
		return client, nil
	}

	config, err := buildRestConfig(kubeConfigPath, clusterName)
	if err != nil {
		return nil, fmt.Errorf("error building config: %s", err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error building clientSet: %s", err)
	}

	cm.clients[clusterName] = clientSet

	return clientSet, nil
}

func (cm *ClusterManager) GetDynamicClient(clusterName string, kubeConfigPath string) (*dynamic.DynamicClient, error) {
	cm.mu.RLock()
	if dynamicClient, exists := cm.dynamicClients[clusterName]; exists {
		cm.mu.RUnlock()
		return dynamicClient, nil
	}
	cm.mu.RUnlock()

	cm.mu.Lock()
	defer cm.mu.Unlock()

	if dynamicClient, exists := cm.dynamicClients[clusterName]; exists {
		return dynamicClient, nil
	}

	config, err := buildRestConfig(kubeConfigPath, clusterName)
	if err != nil {
		return nil, fmt.Errorf("error building config: %s", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error building dynamicClient: %s", err)
	}

	cm.dynamicClients[clusterName] = dynamicClient

	return dynamicClient, nil
}

func (cm *ClusterManager) GetMetricClient(clusterName string, conf *rest.Config) (*metricsv.Clientset, error) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if client, exists := cm.metricClients[clusterName]; exists {
		return client, nil
	}

	metricClient, err := metricsv.NewForConfig(conf)
	if err != nil {
		return nil, fmt.Errorf("error building metric client: %w", err)
	}

	cm.metricClients[clusterName] = metricClient

	return metricClient, nil
}

// TODO: review this code and refactor
func buildRestConfig(kubeConfigPath, context string) (*rest.Config, error) {
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: context}
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	return clientConfig.ClientConfig()
}
