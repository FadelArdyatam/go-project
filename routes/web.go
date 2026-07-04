package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/support"

	"cloud-compute/app/facades"
	"cloud-compute/app/http/controllers"
	appmiddleware "cloud-compute/app/http/middleware"
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

	authController := controllers.NewAuthController()
	facades.Route().Get("/register", authController.ShowRegister)
	facades.Route().Post("/register", authController.Register)
	facades.Route().Get("/login", authController.ShowLogin)
	facades.Route().Post("/login", authController.Login)
	facades.Route().Post("/logout", authController.Logout)

	noteController := controllers.NewNoteController()
	facades.Route().Middleware(appmiddleware.Auth()).Group(func(router route.Router) {
		router.Get("/notes", noteController.Index)
		router.Get("/notes/create", noteController.ShowCreate)
		router.Post("/notes", noteController.Store)
		router.Get("/notes/{id}/edit", noteController.ShowEdit)
		router.Post("/notes/{id}", noteController.Update)
		router.Post("/notes/{id}/delete", noteController.Destroy)
	})
}
