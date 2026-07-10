package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ibrohimubarok/task-tracker-api/internal/api"
	"github.com/ibrohimubarok/task-tracker-api/internal/config"
	"github.com/ibrohimubarok/task-tracker-api/internal/database"
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
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	api.Router(r)
	log.Printf("Server running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
