package usecase

import (
	"Kubexplorer/internal/model"
	"context"
)

type StorageUseCase interface {
	GetPersistentVolumes(ctx context.Context, clusterCtx string) ([]model.PersistentVolumeDto, error)
	GetPersistentVolume(ctx context.Context, ref model.ResourceRef) (model.PersistentVolumeDto, error)
	UpdatePersistentVolume(ctx context.Context, ref model.ResourceRef, dto model.PersistentVolumeDto) error
	DeletePersistentVolume(ctx context.Context, ref model.ResourceRef) error

	GetPersistentVolumesClaim(ctx context.Context, clusterCtx string) ([]model.PersistentVolumeClaimDto, error)
	GetPersistentVolumeClaim(ctx context.Context, ref model.ResourceRef) (model.PersistentVolumeClaimDto, error)
	UpdatePersistentVolumeClaim(ctx context.Context, ref model.ResourceRef, dto model.PersistentVolumeClaimDto) error
	DeletePersistentVolumeClaim(ctx context.Context, ref model.ResourceRef) error
}

type storageUseCase struct {
	client StorageClient
}

func NewStorageUseCase(client StorageClient) StorageUseCase {
	return &storageUseCase{client: client}
}

func (s *storageUseCase) GetPersistentVolumes(ctx context.Context, clusterCtx string) ([]model.PersistentVolumeDto, error) {
	return s.client.GetPersistentVolumes(ctx, clusterCtx)
}

func (s *storageUseCase) GetPersistentVolume(ctx context.Context, ref model.ResourceRef) (model.PersistentVolumeDto, error) {
	return s.client.GetPersistentVolume(ctx, ref)
}

func (s *storageUseCase) UpdatePersistentVolume(ctx context.Context, ref model.ResourceRef, dto model.PersistentVolumeDto) error {
	return s.client.UpdatePersistentVolume(ctx, ref, dto)
}

func (s *storageUseCase) DeletePersistentVolume(ctx context.Context, ref model.ResourceRef) error {
	return s.client.DeletePersistentVolume(ctx, ref)
}

func (s *storageUseCase) GetPersistentVolumesClaim(ctx context.Context, clusterCtx string) ([]model.PersistentVolumeClaimDto, error) {
	return s.client.GetPersistentVolumesClaim(ctx, clusterCtx)
}

func (s *storageUseCase) GetPersistentVolumeClaim(ctx context.Context, ref model.ResourceRef) (model.PersistentVolumeClaimDto, error) {
	return s.client.GetPersistentVolumeClaim(ctx, ref)
}

func (s *storageUseCase) UpdatePersistentVolumeClaim(ctx context.Context, ref model.ResourceRef, dto model.PersistentVolumeClaimDto) error {
	return s.client.UpdatePersistentVolumeClaim(ctx, ref, dto)
}

func (s *storageUseCase) DeletePersistentVolumeClaim(ctx context.Context, ref model.ResourceRef) error {
	return s.client.DeletePersistentVolumeClaim(ctx, ref)
}
