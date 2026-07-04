package controllers

import (
	"github.com/goravel/framework/contracts/http"

	"cloud-compute/app/facades"
	"cloud-compute/app/models"
)

type NoteController struct{}

func NewNoteController() *NoteController {
	return &NoteController{}
}

func currentUser(ctx http.Context) models.User {
	var user models.User
	_ = facades.Auth(ctx).User(&user)
	return user
}

func (r *NoteController) Index(ctx http.Context) http.Response {
	user := currentUser(ctx)

	var notes []models.Note
	_ = facades.Orm().Query().Where("user_id = ?", user.ID).OrderByDesc("id").Get(&notes)

	return ctx.Response().View().Make("notes_index.tmpl", map[string]any{
		"user":  user,
		"notes": notes,
	})
}

func (r *NoteController) ShowCreate(ctx http.Context) http.Response {
	return ctx.Response().View().Make("notes_form.tmpl", map[string]any{
		"action": "/notes",
		"note":   models.Note{},
	})
}

func (r *NoteController) Store(ctx http.Context) http.Response {
	note := models.Note{
		UserID: currentUser(ctx).ID,
		Title:  ctx.Request().Input("title"),
		Body:   ctx.Request().Input("body"),
	}

	_ = facades.Orm().Query().Create(&note)

	return ctx.Response().Redirect(http.StatusFound, "/notes")
}

func (r *NoteController) ShowEdit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	var note models.Note
	if err := facades.Orm().Query().Where("id = ? AND user_id = ?", id, currentUser(ctx).ID).First(&note); err != nil || note.ID == 0 {
		return ctx.Response().Redirect(http.StatusFound, "/notes")
	}

	return ctx.Response().View().Make("notes_form.tmpl", map[string]any{
		"action": "/notes/" + id,
		"note":   note,
	})
}

func (r *NoteController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	var note models.Note
	if err := facades.Orm().Query().Where("id = ? AND user_id = ?", id, currentUser(ctx).ID).First(&note); err != nil || note.ID == 0 {
		return ctx.Response().Redirect(http.StatusFound, "/notes")
	}

	note.Title = ctx.Request().Input("title")
	note.Body = ctx.Request().Input("body")
	_ = facades.Orm().Query().Save(&note)

	return ctx.Response().Redirect(http.StatusFound, "/notes")
}

func (r *NoteController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	_, _ = facades.Orm().Query().Where("id = ? AND user_id = ?", id, currentUser(ctx).ID).Delete(&models.Note{})

	return ctx.Response().Redirect(http.StatusFound, "/notes")
}
