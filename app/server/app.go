package server

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"key-value-storage/app/middleware"
	"net/http"
)

type app struct {}

func NewApp() *app {
	return &app{}
}

var port = flag.Uint64("p", 8000, "port")

func (a *app) Start() {
	flag.Parse()

	router := mux.NewRouter()
	router.Use(middleware.RecoveryMiddleware)

	http.Handle("/", router)

	golog.Infof("Server tarted at port %d", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		golog.Error("Server failed: ", err.Error())
	}
}
