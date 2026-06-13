package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StorageClient struct {
	manager ClusterResolver
}

func NewStorageClient(manager ClusterResolver) *StorageClient {
	return &StorageClient{manager: manager}
}

func (s StorageClient) GetPersistentVolumes(ctx context.Context, clusterCtx string) ([]model.PersistentVolumeDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	volumes, err := client.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("listing persistent volumes: %w", err)
	}

	var result []model.PersistentVolumeDto

	for _, volume := range volumes.Items {
		result = append(result, mapPersistentVolume(volume))
	}

	return result, nil
}

func (s StorageClient) GetPersistentVolume(ctx context.Context, ref model.ResourceRef) (model.PersistentVolumeDto, error) {
	name, clusterCtx := ref.Name, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.PersistentVolumeDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	volume, err := client.CoreV1().PersistentVolumes().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return model.PersistentVolumeDto{}, fmt.Errorf("getting persistent volume %s: %w", name, err)
	}

	return mapPersistentVolume(*volume), nil
}

func (s StorageClient) UpdatePersistentVolume(ctx context.Context, ref model.ResourceRef, dto model.PersistentVolumeDto) error {
	name, clusterCtx := ref.Name, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	c := client.CoreV1().PersistentVolumes()
	volume, err := c.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("getting persistent volume %s: %w", name, err)
	}

	volume.Name = dto.Name
	volume.Namespace = dto.Namespace

	_, err = c.Update(ctx, volume, metav1.UpdateOptions{})

	return err
}

func (s StorageClient) DeletePersistentVolume(ctx context.Context, ref model.ResourceRef) error {
	name, clusterCtx := ref.Name, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	return client.CoreV1().PersistentVolumes().Delete(ctx, name, metav1.DeleteOptions{})
}

func (s StorageClient) GetPersistentVolumesClaim(ctx context.Context, clusterCtx string) ([]model.PersistentVolumeClaimDto, error) {
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	volumeClaims, err := client.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("listing persistent volume claims: %w", err)
	}

	var result []model.PersistentVolumeClaimDto

	for _, volumeClaim := range volumeClaims.Items {
		result = append(result, mapPersistentVolumeClaim(volumeClaim))
	}

	return result, nil
}

func (s StorageClient) GetPersistentVolumeClaim(ctx context.Context, ref model.ResourceRef) (model.PersistentVolumeClaimDto, error) {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return model.PersistentVolumeClaimDto{}, fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	volumeClaim, err := client.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return model.PersistentVolumeClaimDto{}, fmt.Errorf("getting persistent volume claim %s/%s: %w", namespace, name, err)
	}

	return mapPersistentVolumeClaim(*volumeClaim), nil
}

func (s StorageClient) UpdatePersistentVolumeClaim(ctx context.Context, ref model.ResourceRef, dto model.PersistentVolumeClaimDto) error {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	c := client.CoreV1().PersistentVolumeClaims(namespace)
	volumeClaim, err := c.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("getting persistent volume claim %s/%s: %w", namespace, name, err)
	}

	volumeClaim.Name = dto.Name
	volumeClaim.Namespace = dto.Namespace

	_, err = c.Update(ctx, volumeClaim, metav1.UpdateOptions{})

	return err
}

func (s StorageClient) DeletePersistentVolumeClaim(ctx context.Context, ref model.ResourceRef) error {
	name, namespace, clusterCtx := ref.Name, ref.Namespace, ref.Cluster
	client, err := s.manager.ResolveClusterContext(clusterCtx)
	if err != nil {
		return fmt.Errorf("kubeclient: error resolving cluster context: %w", err)
	}

	return client.CoreV1().PersistentVolumeClaims(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func mapPersistentVolume(volume v1.PersistentVolume) model.PersistentVolumeDto {
	var stringAccessModes []string
	for _, mode := range volume.Spec.AccessModes {
		stringAccessModes = append(stringAccessModes, string(mode))
	}

	spec := model.VolumeSpec{
		VolumeMode:                    volumeModeString(volume.Spec.VolumeMode),
		AccessModes:                   stringAccessModes,
		StorageClass:                  volume.Spec.StorageClassName,
		PersistentVolumeReclaimPolicy: string(volume.Spec.PersistentVolumeReclaimPolicy),
		MountOptions:                  volume.Spec.MountOptions,
		Capacity:                      mapResourceList(volume.Spec.Capacity),
	}

	if local := volume.Spec.Local; local != nil {
		spec.Local.Path = local.Path
		if local.FSType != nil {
			spec.Local.FSType = *local.FSType
		}
	}

	if hostPath := volume.Spec.HostPath; hostPath != nil {
		spec.Host.Path = hostPath.Path
		if hostPath.Type != nil {
			spec.Host.Type = string(*hostPath.Type)
		}
	}

	if nfs := volume.Spec.NFS; nfs != nil {
		spec.NFS.Server = nfs.Server
		spec.NFS.Path = nfs.Path
	}

	claim := ""
	if volume.Spec.ClaimRef != nil {
		claim = volume.Spec.ClaimRef.Name
	}

	return model.PersistentVolumeDto{
		Name:         volume.Name,
		Namespace:    volume.Namespace,
		StorageClass: volume.Spec.StorageClassName,
		Capacity:     volume.Spec.Capacity.Storage().String(),
		Claim:        claim,
		Labels:       volume.Labels,
		Status:       string(volume.Status.Phase),
		Age:          model.FormatAge(volume.CreationTimestamp.Time),
		CreatedAt:    volume.CreationTimestamp.Unix(),
		VolumeSpec:   spec,
	}
}

func mapPersistentVolumeClaim(volumeClaim v1.PersistentVolumeClaim) model.PersistentVolumeClaimDto {
	var stringAccessModes []string
	for _, mode := range volumeClaim.Spec.AccessModes {
		stringAccessModes = append(stringAccessModes, string(mode))
	}

	spec := model.VolumeClaimSpec{
		VolumeName:  volumeClaim.Spec.VolumeName,
		VolumeMode:  volumeModeString(volumeClaim.Spec.VolumeMode),
		AccessModes: stringAccessModes,
		Limit:       mapResourceList(volumeClaim.Spec.Resources.Limits),
		Request:     mapResourceList(volumeClaim.Spec.Resources.Requests),
	}

	if volumeClaim.Spec.DataSource != nil {
		spec.DataSourceName = volumeClaim.Spec.DataSource.Name
	}
	if volumeClaim.Spec.StorageClassName != nil {
		spec.StorageClass = *volumeClaim.Spec.StorageClassName
	}
	if volumeClaim.Spec.VolumeAttributesClassName != nil {
		spec.VolumeAttributesClassName = *volumeClaim.Spec.VolumeAttributesClassName
	}

	return model.PersistentVolumeClaimDto{
		Name:            volumeClaim.Name,
		Namespace:       volumeClaim.Namespace,
		StorageClass:    spec.StorageClass,
		Size:            volumeClaim.Spec.Resources.Requests.Storage().String(),
		Labels:          volumeClaim.Labels,
		Status:          string(volumeClaim.Status.Phase),
		Age:             model.FormatAge(volumeClaim.CreationTimestamp.Time),
		CreatedAt:       volumeClaim.CreationTimestamp.Unix(),
		VolumeClaimSpec: spec,
		Capacity:        mapResourceList(volumeClaim.Status.Capacity),
	}
}

func volumeModeString(mode *v1.PersistentVolumeMode) string {
	if mode == nil {
		return ""
	}
	return string(*mode)
}

func mapResourceList(list v1.ResourceList) model.Resource {
	return model.Resource{
		Cpu:              list.Cpu().MilliValue(),
		Memory:           list.Memory().ScaledValue(resource.Mega),
		Storage:          list.Storage().ScaledValue(resource.Mega),
		StorageEphemeral: list.StorageEphemeral().ScaledValue(resource.Mega),
	}
}
