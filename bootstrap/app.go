package bootstrap

import (
	"github.com/goravel/framework/contracts/foundation/configuration"
	contractsfoundation "github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/foundation"
	sessionmiddleware "github.com/goravel/framework/session/middleware"

	"cloud-compute/config"
	"cloud-compute/routes"
)

func Boot() contractsfoundation.Application {
	return foundation.Setup().
		WithMigrations(Migrations).
		WithRouting(func() {
			routes.Web()
			routes.Grpc()
		}).
		WithProviders(Providers).
		WithConfig(config.Boot).
		WithMiddleware(func(handler configuration.Middleware) {
			handler.Append(sessionmiddleware.StartSession())
		}).
		Create()
}
