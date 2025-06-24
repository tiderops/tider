package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type StorageEndpoint struct {
	useCase usecase.StorageUseCase
}

func NewStorageEndpoint(useCase usecase.StorageUseCase) *StorageEndpoint {
	return &StorageEndpoint{useCase: useCase}
}

func (se *StorageEndpoint) GetPersistentVolumes(clusterCtx string) ([]model.PersistentVolumeDto, error) {
	return se.useCase.GetPersistentVolumes(clusterCtx)
}

func (se *StorageEndpoint) GetPersistentVolume(name string, clusterCtx string) (model.PersistentVolumeDto, error) {
	return se.useCase.GetPersistentVolume(name, clusterCtx)
}

func (se *StorageEndpoint) UpdatePersistentVolume(name string, dto model.PersistentVolumeDto, clusterCtx string) error {
	return se.useCase.UpdatePersistentVolume(name, dto, clusterCtx)
}

func (se *StorageEndpoint) DeletePersistentVolume(name string, clusterCtx string) error {
	return se.useCase.DeletePersistentVolume(name, clusterCtx)
}

func (se *StorageEndpoint) GetPersistentVolumesClaim(clusterCtx string) ([]model.PersistentVolumeClaimDto, error) {
	return se.useCase.GetPersistentVolumesClaim(clusterCtx)
}

func (se *StorageEndpoint) GetPersistentVolumeClaim(name string, namespace string, clusterCtx string) (model.PersistentVolumeClaimDto, error) {
	return se.useCase.GetPersistentVolumeClaim(name, namespace, clusterCtx)
}

func (se *StorageEndpoint) UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto, clusterCtx string) error {
	return se.useCase.UpdatePersistentVolumeClaim(name, namespace, dto, clusterCtx)
}

func (se *StorageEndpoint) DeletePersistentVolumeClaim(name string, namespace string, clusterCtx string) error {
	return se.useCase.DeletePersistentVolumeClaim(name, namespace, clusterCtx)
}
