package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ibrohimubarok/task-tracker/internal/api/router"
	"github.com/ibrohimubarok/task-tracker/internal/app"
	"github.com/ibrohimubarok/task-tracker/internal/config"
	"github.com/ibrohimubarok/task-tracker/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database connected successfully")

	// Dependency Injection
	app := app.New(db, cfg.TokenAuth)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	router.Route(r, app.Handlers, cfg.TokenAuth)
	log.Printf("Server running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
