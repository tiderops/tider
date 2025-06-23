package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type storageClient struct {
	manager ClusterResolver
}

func NewStorageClient(manager ClusterResolver) StorageClient {
	return &storageClient{manager: manager}
}

func (s storageClient) GetPersistentVolumes(clusterCtx string) ([]model.PersistentVolumeDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	volumes, err := client.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var result []model.PersistentVolumeDto

	for _, volume := range volumes.Items {
		var stringAccessModes []string
		for _, mode := range volume.Spec.AccessModes {
			stringAccessModes = append(stringAccessModes, string(mode))
		}

		dto := model.PersistentVolumeDto{
			Name:              volume.Name,
			Namespace:         volume.Namespace,
			CreationTimestamp: volume.CreationTimestamp.String(),
			Labels:            volume.Labels,
			VolumeSpec: model.VolumeSpec{
				Local: model.Local{
					Path:   volume.Spec.Local.Path,
					FSType: *volume.Spec.Local.FSType,
				},
				VolumeMode:                    string(*volume.Spec.VolumeMode),
				AccessModes:                   stringAccessModes,
				StorageClass:                  volume.Spec.StorageClassName,
				VolumeAttributesClassName:     *volume.Spec.VolumeAttributesClassName,
				PersistentVolumeReclaimPolicy: string(volume.Spec.PersistentVolumeReclaimPolicy),
				MountOptions:                  volume.Spec.MountOptions,
				Capacity: model.Resource{
					Cpu:              volume.Spec.Capacity.Cpu().String(),
					Memory:           volume.Spec.Capacity.Memory().String(),
					Storage:          volume.Spec.Capacity.Storage().String(),
					StorageEphemeral: volume.Spec.Capacity.StorageEphemeral().String(),
				},
				Host: model.Host{
					Path: volume.Spec.HostPath.Path,
					Type: string(*volume.Spec.HostPath.Type),
				},
				NFS: model.NFS{
					Server: volume.Spec.NFS.Server,
					Path:   volume.Spec.NFS.Path,
				},
			},
		}

		result = append(result, dto)
	}

	return result, errors.New("Error to get persistent volumes claim")
}

func (s storageClient) GetPersistentVolume(name string, clusterCtx string) (model.PersistentVolumeDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.PersistentVolumeDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	volume, err := client.CoreV1().PersistentVolumes().Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error to get persistent volume")
	}

	var stringAccessModes []string
	for _, mode := range volume.Spec.AccessModes {
		stringAccessModes = append(stringAccessModes, string(mode))
	}

	return model.PersistentVolumeDto{
		Name:              volume.Name,
		Namespace:         volume.Namespace,
		CreationTimestamp: volume.CreationTimestamp.String(),
		Labels:            volume.Labels,
		VolumeSpec: model.VolumeSpec{
			Local: model.Local{
				Path:   volume.Spec.Local.Path,
				FSType: *volume.Spec.Local.FSType,
			},
			VolumeMode:                    string(*volume.Spec.VolumeMode),
			AccessModes:                   stringAccessModes,
			StorageClass:                  volume.Spec.StorageClassName,
			VolumeAttributesClassName:     *volume.Spec.VolumeAttributesClassName,
			PersistentVolumeReclaimPolicy: string(volume.Spec.PersistentVolumeReclaimPolicy),
			MountOptions:                  volume.Spec.MountOptions,
			Capacity: model.Resource{
				Cpu:              volume.Spec.Capacity.Cpu().String(),
				Memory:           volume.Spec.Capacity.Memory().String(),
				Storage:          volume.Spec.Capacity.Storage().String(),
				StorageEphemeral: volume.Spec.Capacity.StorageEphemeral().String(),
			},
			Host: model.Host{
				Path: volume.Spec.HostPath.Path,
				Type: string(*volume.Spec.HostPath.Type),
			},
			NFS: model.NFS{
				Server: volume.Spec.NFS.Server,
				Path:   volume.Spec.NFS.Path,
			},
		},
	}, nil
}

func (s storageClient) UpdatePersistentVolume(name string, dto model.PersistentVolumeDto, clusterCtx string) error {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	c := client.CoreV1().PersistentVolumes()
	volume, err := c.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	volume.Name = dto.Name
	volume.Namespace = dto.Namespace

	_, err = c.Update(context.TODO(), volume, metav1.UpdateOptions{})

	return err
}

func (s storageClient) DeletePersistentVolume(name string, clusterCtx string) error {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	return client.CoreV1().PersistentVolumes().Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (s storageClient) GetPersistentVolumesClaim(clusterCtx string) ([]model.PersistentVolumeClaimDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	volumeClaims, err := client.CoreV1().PersistentVolumeClaims("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var result []model.PersistentVolumeClaimDto

	for _, volumeClaim := range volumeClaims.Items {
		var stringAccessModes []string
		for _, mode := range volumeClaim.Spec.AccessModes {
			stringAccessModes = append(stringAccessModes, string(mode))
		}

		dto := model.PersistentVolumeClaimDto{
			Name:              volumeClaim.Name,
			Namespace:         volumeClaim.Namespace,
			CreationTimestamp: volumeClaim.CreationTimestamp.String(),
			Labels:            volumeClaim.Labels,
			VolumeClaimSpec: model.VolumeClaimSpec{
				VolumeName:                volumeClaim.Spec.VolumeName,
				VolumeMode:                string(*volumeClaim.Spec.VolumeMode),
				AccessModes:               stringAccessModes,
				DataSourceName:            volumeClaim.Spec.DataSource.Name,
				StorageClass:              *volumeClaim.Spec.StorageClassName,
				VolumeAttributesClassName: *volumeClaim.Spec.VolumeAttributesClassName,
				Limit: model.Resource{
					Cpu:              volumeClaim.Spec.Resources.Limits.Cpu().String(),
					Memory:           volumeClaim.Spec.Resources.Limits.Memory().String(),
					Storage:          volumeClaim.Spec.Resources.Limits.Storage().String(),
					StorageEphemeral: volumeClaim.Spec.Resources.Limits.StorageEphemeral().String(),
				},
			},
		}

		result = append(result, dto)
	}

	return result, errors.New("Error to get persistent volumes claim")
}

func (s storageClient) GetPersistentVolumeClaim(name string, namespace string, clusterCtx string) (model.PersistentVolumeClaimDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.PersistentVolumeClaimDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	volumeClaim, err := client.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error to get persistent volume claim")
	}

	var stringAccessModes []string
	for _, mode := range volumeClaim.Spec.AccessModes {
		stringAccessModes = append(stringAccessModes, string(mode))
	}

	return model.PersistentVolumeClaimDto{
		Name:              volumeClaim.Name,
		Namespace:         volumeClaim.Namespace,
		CreationTimestamp: volumeClaim.CreationTimestamp.String(),
		Labels:            volumeClaim.Labels,
		VolumeClaimSpec: model.VolumeClaimSpec{
			VolumeName:                volumeClaim.Spec.VolumeName,
			VolumeMode:                string(*volumeClaim.Spec.VolumeMode),
			AccessModes:               stringAccessModes,
			DataSourceName:            volumeClaim.Spec.DataSource.Name,
			StorageClass:              *volumeClaim.Spec.StorageClassName,
			VolumeAttributesClassName: *volumeClaim.Spec.VolumeAttributesClassName,
			Limit: model.Resource{
				Cpu:              volumeClaim.Spec.Resources.Limits.Cpu().String(),
				Memory:           volumeClaim.Spec.Resources.Limits.Memory().String(),
				Storage:          volumeClaim.Spec.Resources.Limits.Storage().String(),
				StorageEphemeral: volumeClaim.Spec.Resources.Limits.StorageEphemeral().String(),
			},
		},
	}, errors.New("Error to get persistent volumes claim")
}

func (s storageClient) UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto, clusterCtx string) error {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	c := client.CoreV1().PersistentVolumeClaims(namespace)
	volumeClaim, err := c.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	volumeClaim.Name = dto.Name
	volumeClaim.Namespace = dto.Namespace

	_, err = c.Update(context.TODO(), volumeClaim, metav1.UpdateOptions{})

	return err
}

func (s storageClient) DeletePersistentVolumeClaim(name string, namespace string, clusterCtx string) error {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	return client.CoreV1().PersistentVolumeClaims(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}
