package usecase

import storageInterfaces "key-value-storage/app/storage/interfaces"

type storageUseCase struct {
	repository storageInterfaces.StorageRepository
}

func NewStorageUseCase(repository storageInterfaces.StorageRepository) *storageUseCase {
	return &storageUseCase{repository: repository}
}
