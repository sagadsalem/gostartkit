package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sajadsalem/startkit/app"
	"github.com/sajadsalem/startkit/internal/server"
)

func main() {
	app, err := app.NewApp(":8080")
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	server.ListenAndServe(app)
}
