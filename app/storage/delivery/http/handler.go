package storageHttp

import storageInterfaces "key-value-storage/app/storage/interfaces"

type storageHandler struct {
	useCase storageInterfaces.StorageUseCase
}

func NewStorageHandler(useCase storageInterfaces.StorageUseCase) *storageHandler {
	return &storageHandler{useCase: useCase}
}
