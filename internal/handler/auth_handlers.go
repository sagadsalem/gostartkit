package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomeHandler is a handler function for the home route
func (h *Handler) HomeHandler(c echo.Context) error {
	return c.String(http.StatusBadRequest, "Hello Gophers!")
}
