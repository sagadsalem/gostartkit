package handler

import (
	"github.com/sajadsalem/startkit/app"
)

type Handler struct {
	*app.App
}

type Error struct {
	Message          string              `json:"message"`
	ValidationErrors []map[string]string `json:"validationErrors,omitempty"`
}

type Success struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewHandler(app *app.App) *Handler {
	return &Handler{app}
}
