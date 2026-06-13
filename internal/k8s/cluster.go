package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ClusterClient struct {
	manager *ClusterManager
}

func NewCluster(manager *ClusterManager) *ClusterClient {
	return &ClusterClient{manager: manager}
}

func listKubeConfigFiles() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	kubeDir := filepath.Join(home, ".kube")

	files, err := os.ReadDir(kubeDir)
	if err != nil {
		return nil, err
	}

	var configs []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml") || name == "config" {
			configs = append(configs, filepath.Join(kubeDir, name))
		}
	}

	return configs, nil
}

func (c *ClusterClient) ListAvailableClusters(ctx context.Context) ([]model.ClusterInfo, error) {
	files, err := listKubeConfigFiles()
	if err != nil {
		return nil, err
	}

	var clusters []model.ClusterInfo

	for _, file := range files {
		config, err := clientcmd.LoadFromFile(file)
		if err != nil {
			slog.Warn("skipping kubeconfig file", "file", file, "error", err)
			continue
		}

		for name, kubeCtx := range config.Contexts {
			clt := config.Clusters[kubeCtx.Cluster]
			if clt == nil {
				continue
			}

			status := c.checkClusterStatus(ctx, file, name)
			ci := model.ClusterInfo{
				Name:      name,
				Cluster:   kubeCtx.Cluster,
				Server:    clt.Server,
				User:      kubeCtx.AuthInfo,
				Namespace: kubeCtx.Namespace,
				Status:    status,
				Source:    file,
			}
			clusters = append(clusters, ci)
		}
	}

	return clusters, nil
}

func (c *ClusterClient) checkClusterStatus(ctx context.Context, kubeConfigPath string, contextName string) bool {
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: contextName}
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath}
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides)

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return false
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return false
	}

	if _, err := c.manager.GetClient(contextName, kubeConfigPath); err != nil {
		slog.Error("registering cluster client failed", "cluster", contextName, "error", err)
		return false
	}
	if _, err := c.manager.GetDynamicClient(contextName, kubeConfigPath); err != nil {
		slog.Error("registering dynamic client failed", "cluster", contextName, "error", err)
		return false
	}
	if _, err := c.manager.GetMetricClient(contextName, restConfig); err != nil {
		slog.Error("registering metric client failed", "cluster", contextName, "error", err)
		return false
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	done := make(chan bool, 1)
	go func() {
		_, err := clientset.Discovery().ServerVersion()
		done <- err == nil
	}()

	select {
	case ok := <-done:
		return ok
	case <-ctx.Done():
		return false
	}
}

func (c *ClusterClient) GetCurrentCluster(ctx context.Context, name string) (model.EnvironmentDto, error) {
	dto := model.EnvironmentDto{
		Name:        "minikube",
		Description: "minikube description",
		Env:         "Dev",
		Status:      true,
	}

	return dto, nil
}

func (c *ClusterClient) GetClusterObjects(ctx context.Context, clusterCtx string) (model.ObjectMapDto, error) {
	client, err := c.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("resolving cluster %q: %w", clusterCtx, err)
	}

	pods, err := client.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing pods: %w", err)
	}

	rs, err := client.AppsV1().ReplicaSets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing replicasets: %w", err)
	}

	dps, err := client.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing deployments: %w", err)
	}

	svcs, err := client.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing services: %w", err)
	}

	nodes, err := client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing nodes: %w", err)
	}

	pvcs, err := client.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing PVCs: %w", err)
	}

	cms, err := client.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing configmaps: %w", err)
	}

	secrets, err := client.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing secrets: %w", err)
	}

	nss, err := client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return model.ObjectMapDto{}, fmt.Errorf("listing namespaces: %w", err)
	}

	return model.ObjectMapDto{
		Pods:                   model.MapPodsToDto(pods),
		ReplicaSets:            model.MapReplicaSetsToDto(rs),
		Deployments:            model.MapDeploymentsToDto(dps),
		Services:               model.MapServicesToDto(svcs),
		Nodes:                  model.MapNodesToDto(nodes),
		PersistentVolumeClaims: model.MapPVCsToDto(pvcs),
		ConfigMaps:             model.MapConfigMapsToDto(cms),
		Secrets:                model.MapSecretsToDto(secrets),
		Namespaces:             model.MapNamespacesToDto(nss),
	}, nil
}
