package controllers

import (
	"github.com/goravel/framework/contracts/http"

	"cloud-compute/app/facades"
	"cloud-compute/app/models"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (r *AuthController) ShowRegister(ctx http.Context) http.Response {
	return ctx.Response().View().Make("register.tmpl", map[string]any{
		"error": ctx.Request().Query("error"),
	})
}

func (r *AuthController) Register(ctx http.Context) http.Response {
	name := ctx.Request().Input("name")
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	if name == "" || email == "" || password == "" {
		return ctx.Response().Redirect(http.StatusFound, "/register?error=Semua+field+wajib+diisi")
	}

	var existing models.User
	if err := facades.Orm().Query().Where("email = ?", email).First(&existing); err == nil && existing.ID != 0 {
		return ctx.Response().Redirect(http.StatusFound, "/register?error=Email+sudah+terdaftar")
	}

	hashed, err := facades.Hash().Make(password)
	if err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/register?error=Gagal+memproses+password")
	}

	user := models.User{Name: name, Email: email, Password: hashed}
	if err := facades.Orm().Query().Create(&user); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/register?error=Gagal+membuat+akun")
	}

	if _, err := facades.Auth(ctx).Login(&user); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/login")
	}

	return ctx.Response().Redirect(http.StatusFound, "/notes")
}

func (r *AuthController) ShowLogin(ctx http.Context) http.Response {
	return ctx.Response().View().Make("login.tmpl", map[string]any{
		"error": ctx.Request().Query("error"),
	})
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	var user models.User
	if err := facades.Orm().Query().Where("email = ?", email).First(&user); err != nil || user.ID == 0 {
		return ctx.Response().Redirect(http.StatusFound, "/login?error=Email+atau+password+salah")
	}

	if !facades.Hash().Check(password, user.Password) {
		return ctx.Response().Redirect(http.StatusFound, "/login?error=Email+atau+password+salah")
	}

	if _, err := facades.Auth(ctx).Login(&user); err != nil {
		return ctx.Response().Redirect(http.StatusFound, "/login?error=Gagal+login")
	}

	return ctx.Response().Redirect(http.StatusFound, "/notes")
}

func (r *AuthController) Logout(ctx http.Context) http.Response {
	_ = facades.Auth(ctx).Logout()

	return ctx.Response().Redirect(http.StatusFound, "/login")
}
