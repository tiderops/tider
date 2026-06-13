package k8s

import (
	"sync"
	"testing"

	"k8s.io/client-go/rest"
)

func TestResolveUnregisteredClusterReturnsError(t *testing.T) {
	cm := NewClusterManager()

	if _, err := cm.ResolveClusterContext("nope"); err == nil {
		t.Error("ResolveClusterContext() expected error for unregistered cluster")
	}
	if _, err := cm.ResolveClusterContextDynamic("nope"); err == nil {
		t.Error("ResolveClusterContextDynamic() expected error for unregistered cluster")
	}
	if _, err := cm.ResolveClusterMetric("nope"); err == nil {
		t.Error("ResolveClusterMetric() expected error for unregistered cluster")
	}
}

// Run with -race: registration and resolution happen concurrently when
// several Wails calls arrive while clusters are still being registered
func TestClusterManagerConcurrentAccess(t *testing.T) {
	cm := NewClusterManager()
	conf := &rest.Config{Host: "http://127.0.0.1:1"}

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			if _, err := cm.GetMetricClient("test-cluster", conf); err != nil {
				t.Errorf("GetMetricClient() unexpected error: %v", err)
			}
		}()
		go func() {
			defer wg.Done()
			_, _ = cm.ResolveClusterMetric("test-cluster")
		}()
	}
	wg.Wait()

	if _, err := cm.ResolveClusterMetric("test-cluster"); err != nil {
		t.Errorf("ResolveClusterMetric() after registration: %v", err)
	}
}
