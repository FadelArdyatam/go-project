package models

import (
	"github.com/goravel/framework/database/orm"
)

type Note struct {
	orm.Model
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
