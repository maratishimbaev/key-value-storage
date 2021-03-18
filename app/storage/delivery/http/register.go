package storageHttp

import (
	"github.com/gorilla/mux"
	storageInterfaces "key-value-storage/app/storage/interfaces"
)

func RegisterHttpEndpoints(router *mux.Router, useCase storageInterfaces.StorageUseCase) {
	h := NewStorageHandler(useCase)

	router.HandleFunc("/", h.CreateOrUpdateKey).Methods("POST")
	router.HandleFunc("/{key:[a-zA-Z]+}", h.DeleteKey).Methods("DELETE")
	router.HandleFunc("/{key:[a-zA-Z]+}", h.GetKeyValue).Methods("GET")
	router.HandleFunc("/", h.GetListOfKeys).Methods("GET")
}
