package server

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sajadsalem/startkit/app"
	"github.com/sajadsalem/startkit/internal/handler"
)

func ListenAndServe(app *app.App) error {
	server := &http.Server{
		Addr:         app.Addr,
		Handler:      routes(app),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Server listening on http://localhost%s", app.Addr)
	return server.ListenAndServe()
}

func routes(app *app.App) http.Handler {
	// create a new handler instance to pass the app instance
	handler := handler.NewHandler(app)
	// create a new echo instance
	e := echo.New()

	// routes
	e.GET("/", handler.HomeHandler)

	return e
}
