package storageHttp

import (
	"github.com/gorilla/mux"
	storageInterfaces "key-value-storage/app/storage/interfaces"
)

func RegisterHttpEndpoints(router *mux.Router, useCase storageInterfaces.StorageUseCase) {
	_ = NewStorageHandler(useCase)
}
