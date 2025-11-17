package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/lrs-rodrigues/estudo-golang/internal/di"
)

func main() {
	var db = di.NewPostregresDB()
	var repos = di.NewRepositories(db)
	var services = di.NewServices(repos)
	var handler = di.NewHandler(services)
	var router = di.NewRouter(handler)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
