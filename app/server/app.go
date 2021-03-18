package server

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"key-value-storage/app/middleware"
	storageHttp "key-value-storage/app/storage/delivery/http"
	storageInterfaces "key-value-storage/app/storage/interfaces"
	storageMap "key-value-storage/app/storage/repository/map"
	"key-value-storage/app/storage/usecase"
	"net/http"
)

type app struct {
	storageUseCase storageInterfaces.StorageUseCase
}

func NewApp() *app {
	storageRepository := storageMap.NewStorageRepository()

	return &app{
		storageUseCase: usecase.NewStorageUseCase(storageRepository),
	}
}

var port = flag.Uint64("p", 8000, "port")

func (a *app) Start() {
	flag.Parse()

	router := mux.NewRouter()
	router.Use(middleware.RecoveryMiddleware)

	storageHttp.RegisterHttpEndpoints(router, a.storageUseCase)

	http.Handle("/", router)

	golog.Infof("Server tarted at port %d", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		golog.Error("Server failed: ", err.Error())
	}
}
