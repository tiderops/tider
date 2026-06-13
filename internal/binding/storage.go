package binding

import (
	"Kubexplorer/internal/apperr"
	"Kubexplorer/internal/k8s"
	"Kubexplorer/internal/model"
	"Kubexplorer/internal/usecase"
)

type Storage struct {
	app     *App
	storage usecase.StorageUseCase
}

func BuildStorage(app *App, manager *k8s.ClusterManager) *Storage {
	return &Storage{
		app:     app,
		storage: usecase.NewStorageUseCase(k8s.NewStorageClient(manager)),
	}
}

func (s *Storage) GetPersistentVolumes(clusterCtx string) ([]model.PersistentVolumeDto, error) {
	ctx, cancel := s.app.requestContext()
	defer cancel()

	volumes, err := s.storage.GetPersistentVolumes(ctx, clusterCtx)
	return volumes, apperr.Normalize(err)
}

func (s *Storage) GetPersistentVolume(name string, clusterCtx string) (model.PersistentVolumeDto, error) {
	ctx, cancel := s.app.requestContext()
	defer cancel()

	volume, err := s.storage.GetPersistentVolume(ctx, model.ResourceRef{Cluster: clusterCtx, Name: name})
	return volume, apperr.Normalize(err)
}

func (s *Storage) UpdatePersistentVolume(name string, dto model.PersistentVolumeDto, clusterCtx string) error {
	ctx, cancel := s.app.requestContext()
	defer cancel()

	return apperr.Normalize(s.storage.UpdatePersistentVolume(ctx, model.ResourceRef{Cluster: clusterCtx, Name: name}, dto))
}

func (s *Storage) DeletePersistentVolume(name string, clusterCtx string) error {
	ctx, cancel := s.app.requestContext()
	defer cancel()

	return apperr.Normalize(s.storage.DeletePersistentVolume(ctx, model.ResourceRef{Cluster: clusterCtx, Name: name}))
}

func (s *Storage) GetPersistentVolumesClaim(clusterCtx string) ([]model.PersistentVolumeClaimDto, error) {
	ctx, cancel := s.app.requestContext()
	defer cancel()

	claims, err := s.storage.GetPersistentVolumesClaim(ctx, clusterCtx)
	return claims, apperr.Normalize(err)
}

func (s *Storage) GetPersistentVolumeClaim(name string, namespace string, clusterCtx string) (model.PersistentVolumeClaimDto, error) {
	ctx, cancel := s.app.requestContext()
	defer cancel()

	claim, err := s.storage.GetPersistentVolumeClaim(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name})
	return claim, apperr.Normalize(err)
}

func (s *Storage) UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto, clusterCtx string) error {
	ctx, cancel := s.app.requestContext()
	defer cancel()

	return apperr.Normalize(s.storage.UpdatePersistentVolumeClaim(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name}, dto))
}

func (s *Storage) DeletePersistentVolumeClaim(name string, namespace string, clusterCtx string) error {
	ctx, cancel := s.app.requestContext()
	defer cancel()

	return apperr.Normalize(s.storage.DeletePersistentVolumeClaim(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name}))
}
