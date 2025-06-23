package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type cluster struct {
	manager ClusterResolver
}

func NewCluster() ClusterClient {
	return &cluster{}
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

func (c *cluster) ListAvailableClusters() ([]model.ClusterInfo, error) {
	files, err := listKubeConfigFiles()
	if err != nil {
		return nil, err
	}

	var clusters []model.ClusterInfo

	for _, file := range files {
		config, err := clientcmd.LoadFromFile(file)
		if err != nil {
			fmt.Printf("Skipping file %s: %v", file, err)
		}

		for name, ctx := range config.Contexts {
			clt := config.Clusters[ctx.Cluster]
			if clt == nil {
				continue
			}

			status := checkClusterStatus(file, name)
			ci := model.ClusterInfo{
				Name:      name,
				Cluster:   ctx.Cluster,
				Server:    clt.Server,
				User:      ctx.AuthInfo,
				Namespace: ctx.Namespace,
				Status:    status,
				Source:    file,
			}
			clusters = append(clusters, ci)
		}
	}

	return clusters, nil
}

func checkClusterStatus(kubeConfigPath string, contextName string) bool {
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

	fmt.Printf("Init cluster register %s, %s \n", contextName, kubeConfigPath)
	_, err = GlobalClusterManager.GetClient(contextName, kubeConfigPath)

	if err != nil {
		fmt.Println("Error register cluster")
	}
	fmt.Println("Clusters already registered")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan bool, 1)
	go func() {
		_, err = clientset.Discovery().ServerVersion()
		done <- err == nil
	}()

	select {
	case ok := <-done:
		return ok
	case <-ctx.Done():
		return false
	}
}

func (c *cluster) GetCurrentCluster(name string) (model.EnvironmentDto, error) {
	dto := model.EnvironmentDto{
		Name:        "minikube",
		Description: "minikube description",
		Env:         "Dev",
		Status:      true,
	}

	return dto, nil
}
