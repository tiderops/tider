package middleware

import (
	"Kubexplorer/backend/endpoint"
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type StorageMiddleware struct {
	endpoint endpoint.StorageEndpoint
}

func NewStorageMiddleware(endpoint *endpoint.StorageEndpoint) *StorageMiddleware {
	return &StorageMiddleware{endpoint: *endpoint}
}

func (s *StorageMiddleware) GetPersistentVolumes(clusterCtx string) ([]model.PersistentVolumeDto, error) {
	return s.endpoint.GetPersistentVolumes(clusterCtx)
}

func (s *StorageMiddleware) GetPersistentVolume(name string, clusterCtx string) (model.PersistentVolumeDto, error) {
	return s.endpoint.GetPersistentVolume(name, clusterCtx)
}

func (s *StorageMiddleware) UpdatePersistentVolume(name string, dto model.PersistentVolumeDto, clusterCtx string) error {
	return s.endpoint.UpdatePersistentVolume(name, dto, clusterCtx)
}

func (s *StorageMiddleware) DeletePersistentVolume(name string, clusterCtx string) error {
	return s.endpoint.DeletePersistentVolume(name, clusterCtx)
}

func (s *StorageMiddleware) GetPersistentVolumesClaim(clusterCtx string) ([]model.PersistentVolumeClaimDto, error) {
	return s.endpoint.GetPersistentVolumesClaim(clusterCtx)
}

func (s *StorageMiddleware) GetPersistentVolumeClaim(name string, namespace string, clusterCtx string) (model.PersistentVolumeClaimDto, error) {
	return s.endpoint.GetPersistentVolumeClaim(name, namespace, clusterCtx)
}

func (s *StorageMiddleware) UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto, clusterCtx string) error {
	return s.endpoint.UpdatePersistentVolumeClaim(name, namespace, dto, clusterCtx)
}

func (s *StorageMiddleware) DeletePersistentVolumeClaim(name string, namespace string, clusterCtx string) error {
	return s.endpoint.DeletePersistentVolumeClaim(name, namespace, clusterCtx)
}

func BuildStorage(manager kubeclient.ClusterResolver) *StorageMiddleware {
	storageClient := kubeclient.NewStorageClient(manager)

	storageUseCase := usecase.NewStorageUseCase(storageClient)

	storageEndpoint := endpoint.NewStorageEndpoint(storageUseCase)

	return NewStorageMiddleware(storageEndpoint)
}
