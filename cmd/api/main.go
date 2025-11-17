package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	adapterhttp "github.com/lrs-rodrigues/estudo-golang/internal/adapter/http"
	"github.com/lrs-rodrigues/estudo-golang/internal/infra/postgres"
	"github.com/lrs-rodrigues/estudo-golang/internal/services"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL not set")

	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	userRepo := postgres.NewUserRepositoryPostgres(db)
	userService := services.NewUserService(userRepo)
	userHandler := adapterhttp.NewUserHandler(userService)

	userHandler.RegisterRoutes(router)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
