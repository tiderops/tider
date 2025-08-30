package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
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
		panic("Failed to list pv")
	}

	var result []model.PersistentVolumeDto
	fmt.Println("volume.Name", volumes.Items[0])
	for _, volume := range volumes.Items {

		var stringAccessModes []string
		for _, mode := range volume.Spec.AccessModes {
			stringAccessModes = append(stringAccessModes, string(mode))
		}

		dto := model.PersistentVolumeDto{
			Name:      volume.Name,
			Namespace: volume.Namespace,
			Labels:    volume.Labels,
			Status:    string(volume.Status.Phase),
			Age:       volume.CreationTimestamp.String(),
			VolumeSpec: model.VolumeSpec{
				Local: model.Local{
					//Path: volume.Spec.Local.Path, TODO: Validate null reference
					//FSType: *volume.Spec.Local.FSType, TODO: Validate null reference
				},
				VolumeMode:   string(*volume.Spec.VolumeMode),
				AccessModes:  stringAccessModes,
				StorageClass: volume.Spec.StorageClassName,
				//VolumeAttributesClassName:     *volume.Spec.VolumeAttributesClassName, TODO: Validate null reference
				PersistentVolumeReclaimPolicy: string(volume.Spec.PersistentVolumeReclaimPolicy),
				MountOptions:                  volume.Spec.MountOptions,
				Capacity: model.Resource{
					Cpu:              volume.Spec.Capacity.Cpu().MilliValue(),
					Memory:           volume.Spec.Capacity.Memory().ScaledValue(resource.Mega),
					Storage:          volume.Spec.Capacity.Storage().ScaledValue(resource.Mega),
					StorageEphemeral: volume.Spec.Capacity.StorageEphemeral().ScaledValue(resource.Mega),
				},
				Host: model.Host{
					Path: volume.Spec.HostPath.Path,
					Type: string(*volume.Spec.HostPath.Type),
				},
				NFS: model.NFS{
					//Server: volume.Spec.NFS.Server, TODO: Valida null reference
					//Path: volume.Spec.NFS.Path, TODO: Valida null reference
				},
			},
		}

		result = append(result, dto)
	}

	return result, nil
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
		Name:      volume.Name,
		Namespace: volume.Namespace,
		Age:       volume.CreationTimestamp.String(),
		Status:    string(volume.Status.Phase),
		Labels:    volume.Labels,
		VolumeSpec: model.VolumeSpec{
			Local: model.Local{
				//Path:   volume.Spec.Local.Path,
				//FSType: *volume.Spec.Local.FSType,
			},
			VolumeMode:   string(*volume.Spec.VolumeMode),
			AccessModes:  stringAccessModes,
			StorageClass: volume.Spec.StorageClassName,
			//VolumeAttributesClassName:     *volume.Spec.VolumeAttributesClassName,
			PersistentVolumeReclaimPolicy: string(volume.Spec.PersistentVolumeReclaimPolicy),
			MountOptions:                  volume.Spec.MountOptions,
			Capacity: model.Resource{
				Cpu:              volume.Spec.Capacity.Cpu().MilliValue(),
				Memory:           volume.Spec.Capacity.Memory().ScaledValue(resource.Mega),
				Storage:          volume.Spec.Capacity.Storage().ScaledValue(resource.Mega),
				StorageEphemeral: volume.Spec.Capacity.StorageEphemeral().ScaledValue(resource.Mega),
			},
			Host: model.Host{
				Path: volume.Spec.HostPath.Path,
				Type: string(*volume.Spec.HostPath.Type),
			},
			NFS: model.NFS{
				//Server: volume.Spec.NFS.Server,
				//Path:   volume.Spec.NFS.Path,
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
			Name:      volumeClaim.Name,
			Namespace: volumeClaim.Namespace,
			Labels:    volumeClaim.Labels,
			Status:    string(volumeClaim.Status.Phase),
			Age:       volumeClaim.CreationTimestamp.String(),
			VolumeClaimSpec: model.VolumeClaimSpec{
				VolumeName:  volumeClaim.Spec.VolumeName,
				VolumeMode:  string(*volumeClaim.Spec.VolumeMode),
				AccessModes: stringAccessModes,
				//DataSourceName: volumeClaim.Spec.DataSource.Name, TODO: Validate null reference
				StorageClass: *volumeClaim.Spec.StorageClassName,
				//VolumeAttributesClassName: *volumeClaim.Spec.VolumeAttributesClassName, TODO: Validate null reference
				//DataSourceRef: volumeClaim.Spec.DataSourceRef.Name, TODO: Validate null reference
				Limit: model.Resource{
					Cpu:              volumeClaim.Spec.Resources.Limits.Cpu().MilliValue(),
					Memory:           volumeClaim.Spec.Resources.Limits.Memory().ScaledValue(resource.Mega),
					Storage:          volumeClaim.Spec.Resources.Limits.Storage().ScaledValue(resource.Mega),
					StorageEphemeral: volumeClaim.Spec.Resources.Limits.StorageEphemeral().ScaledValue(resource.Mega),
				},
				Request: model.Resource{
					Cpu:              volumeClaim.Spec.Resources.Requests.Cpu().MilliValue(),
					Memory:           volumeClaim.Spec.Resources.Requests.Memory().ScaledValue(resource.Mega),
					Storage:          volumeClaim.Spec.Resources.Requests.Storage().ScaledValue(resource.Mega),
					StorageEphemeral: volumeClaim.Spec.Resources.Requests.StorageEphemeral().ScaledValue(resource.Mega),
				},
			},
			Capacity: model.Resource{
				Cpu:              volumeClaim.Status.Capacity.Cpu().MilliValue(),
				Memory:           volumeClaim.Status.Capacity.Memory().ScaledValue(resource.Mega),
				Storage:          volumeClaim.Status.Capacity.Storage().ScaledValue(resource.Mega),
				StorageEphemeral: volumeClaim.Status.Capacity.StorageEphemeral().ScaledValue(resource.Mega),
			},
		}

		result = append(result, dto)
	}

	return result, nil
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
		Name:      volumeClaim.Name,
		Namespace: volumeClaim.Namespace,
		Age:       volumeClaim.CreationTimestamp.String(),
		Status:    string(volumeClaim.Status.Phase),
		Labels:    volumeClaim.Labels,
		VolumeClaimSpec: model.VolumeClaimSpec{
			VolumeName:                volumeClaim.Spec.VolumeName,
			VolumeMode:                string(*volumeClaim.Spec.VolumeMode),
			AccessModes:               stringAccessModes,
			DataSourceName:            volumeClaim.Spec.DataSource.Name,
			StorageClass:              *volumeClaim.Spec.StorageClassName,
			VolumeAttributesClassName: *volumeClaim.Spec.VolumeAttributesClassName,
			Limit: model.Resource{
				Cpu:              volumeClaim.Spec.Resources.Limits.Cpu().MilliValue(),
				Memory:           volumeClaim.Spec.Resources.Limits.Memory().ScaledValue(resource.Mega),
				Storage:          volumeClaim.Spec.Resources.Limits.Storage().ScaledValue(resource.Mega),
				StorageEphemeral: volumeClaim.Spec.Resources.Limits.StorageEphemeral().ScaledValue(resource.Mega),
			},
		},
	}, nil
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
