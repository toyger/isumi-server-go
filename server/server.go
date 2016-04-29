package server

import (
	"net/http"

	"goji.io"
	"goji.io/pat"

	"github.com/toyger/isumi-server/config"
	"github.com/toyger/isumi-server/server/controller"
)

func Serve(conf *config.Config) {

	rootMux := goji.NewMux()
	rootMux.HandleFuncC(pat.Get("/health_check"), controller.HealthCheck)

	http.ListenAndServe(conf.Bind, rootMux)
}