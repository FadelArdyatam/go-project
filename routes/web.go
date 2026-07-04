package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support"

	"cloud-compute/app/http/controllers"
	"cloud-compute/app/facades"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Static("public", "./public")

	facades.Route().Get("/health", controllers.Health)

	userController := controllers.NewUserController()
	facades.Route().Get("/users", userController.Index)
}
