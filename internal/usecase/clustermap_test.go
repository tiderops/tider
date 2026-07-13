package usecase

import (
	"Kubexplorer/internal/model"
	"testing"
)

func TestBuildObjectMapResolvesRelationships(t *testing.T) {
	obj := &model.ObjectMapDto{
		Nodes:      []model.Node{{Name: "node-1"}},
		Namespaces: []model.Namespace{{Name: "default"}},
		Pods: []model.Pod{
			{
				Name:          "web-7d9f-x1",
				Namespace:     "default",
				NodeName:      "node-1",
				Labels:        map[string]string{"app": "web"},
				OwnerKind:     "ReplicaSet",
				OwnerName:     "web-7d9f",
				PVCNames:      []string{"web-data"},
				ConfigMapRefs: []string{"web-config"},
			},
		},
		ReplicaSets: []model.ReplicaSet{
			{Name: "web-7d9f", Namespace: "default", Deployment: "web"},
		},
		Deployments: []model.Deployment{
			{Name: "web", Namespace: "default"},
		},
		Services: []model.Service{
			{Name: "web-svc", Namespace: "default", Selector: map[string]string{"app": "web"}},
			{Name: "other-svc", Namespace: "default", Selector: map[string]string{"app": "other"}},
		},
		PersistentVolumeClaims: []model.PersistentVolumeClaim{
			{Name: "web-data", Namespace: "default"},
		},
		ConfigMaps: []model.ConfigMap{
			{Name: "web-config", Namespace: "default"},
		},
	}

	BuildObjectMap(obj)

	if got := obj.Deployments[0].PodNames; len(got) != 1 || got[0] != "web-7d9f-x1" {
		t.Errorf("deployment pod names = %v, want [web-7d9f-x1]", got)
	}
	if obj.Pods[0].Deployment != "web" {
		t.Errorf("pod deployment = %q, want web", obj.Pods[0].Deployment)
	}
	if got := obj.Nodes[0].PodNames; len(got) != 1 {
		t.Errorf("node pod names = %v, want one entry", got)
	}
	if got := obj.Services[0].PodNames; len(got) != 1 || got[0] != "web-7d9f-x1" {
		t.Errorf("matching service pod names = %v, want [web-7d9f-x1]", got)
	}
	if got := obj.Services[1].PodNames; len(got) != 0 {
		t.Errorf("non-matching service pod names = %v, want empty", got)
	}
	if got := obj.PersistentVolumeClaims[0].UsedByPods; len(got) != 1 {
		t.Errorf("pvc used-by = %v, want one entry", got)
	}
	if got := obj.ConfigMaps[0].UsedByPods; len(got) != 1 {
		t.Errorf("configmap used-by = %v, want one entry", got)
	}
}

func TestSelectorLabelsMatcher(t *testing.T) {
	tests := []struct {
		name     string
		selector map[string]string
		labels   map[string]string
		want     bool
	}{
		{name: "empty selector never matches", selector: nil, labels: map[string]string{"app": "x"}, want: false},
		{name: "exact match", selector: map[string]string{"app": "x"}, labels: map[string]string{"app": "x"}, want: true},
		{name: "subset match", selector: map[string]string{"app": "x"}, labels: map[string]string{"app": "x", "tier": "web"}, want: true},
		{name: "value mismatch", selector: map[string]string{"app": "x"}, labels: map[string]string{"app": "y"}, want: false},
		{name: "missing key", selector: map[string]string{"app": "x", "tier": "db"}, labels: map[string]string{"app": "x"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectorLabelsMatcher(tt.selector, tt.labels); got != tt.want {
				t.Errorf("selectorLabelsMatcher(%v, %v) = %v, want %v", tt.selector, tt.labels, got, tt.want)
			}
		})
	}
}
