package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type StorageUseCase interface {
	GetPersistentVolumes(clusterCtx string) ([]model.PersistentVolumeDto, error)
	GetPersistentVolume(name string, clusterCtx string) (model.PersistentVolumeDto, error)
	UpdatePersistentVolume(name string, dto model.PersistentVolumeDto, clusterCtx string) error
	DeletePersistentVolume(name string, clusterCtx string) error

	GetPersistentVolumesClaim(clusterCtx string) ([]model.PersistentVolumeClaimDto, error)
	GetPersistentVolumeClaim(name string, namespace string, clusterCtx string) (model.PersistentVolumeClaimDto, error)
	UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto, clusterCtx string) error
	DeletePersistentVolumeClaim(name string, namespace string, clusterCtx string) error
}

type storageUseCase struct {
	client kubeclient.StorageClient
}

func NewStorageUseCase(client kubeclient.StorageClient) StorageUseCase {
	return &storageUseCase{client: client}
}

func (s *storageUseCase) GetPersistentVolumes(clusterCtx string) ([]model.PersistentVolumeDto, error) {
	return s.client.GetPersistentVolumes(clusterCtx)
}

func (s *storageUseCase) GetPersistentVolume(name string, clusterCtx string) (model.PersistentVolumeDto, error) {
	return s.client.GetPersistentVolume(name, clusterCtx)
}

func (s *storageUseCase) UpdatePersistentVolume(name string, dto model.PersistentVolumeDto, clusterCtx string) error {
	return s.client.UpdatePersistentVolume(name, dto, clusterCtx)
}

func (s *storageUseCase) DeletePersistentVolume(name string, clusterCtx string) error {
	return s.client.DeletePersistentVolume(name, clusterCtx)
}

func (s *storageUseCase) GetPersistentVolumesClaim(clusterCtx string) ([]model.PersistentVolumeClaimDto, error) {
	return s.client.GetPersistentVolumesClaim(clusterCtx)
}

func (s *storageUseCase) GetPersistentVolumeClaim(name string, namespace string, clusterCtx string) (model.PersistentVolumeClaimDto, error) {
	return s.client.GetPersistentVolumeClaim(name, namespace, clusterCtx)
}

func (s *storageUseCase) UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto, clusterCtx string) error {
	return s.client.UpdatePersistentVolumeClaim(name, namespace, dto, clusterCtx)
}

func (s *storageUseCase) DeletePersistentVolumeClaim(name string, namespace string, clusterCtx string) error {
	return s.client.DeletePersistentVolumeClaim(name, namespace, clusterCtx)
}
