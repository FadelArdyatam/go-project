package middleware

import (
	"github.com/goravel/framework/contracts/http"

	"cloud-compute/app/facades"
)

func Auth() http.Middleware {
	return func(ctx http.Context) {
		if facades.Auth(ctx).Guard("user").Guest() {
			ctx.Response().Redirect(http.StatusFound, "/login").Abort()
			return
		}

		ctx.Request().Next()
	}
}
