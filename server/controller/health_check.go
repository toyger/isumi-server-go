package controller

import (
	"net/http"

	"golang.org/x/net/context"
	"github.com/toyger/isumi-server/server/render"
)

func HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"status": "ok",
	}
	render.RenderOk(data, w, r)
}
