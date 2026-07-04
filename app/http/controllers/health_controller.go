package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

func Health(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"status": "healthy",
	})
}
