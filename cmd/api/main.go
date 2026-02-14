package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
    router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"status":"ok"}`))
		if err != nil {
			log.Printf("failed to write health response: %v", err)
		}
	})
	router.Get("/doc/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "internal/openapi/openapi.yaml")
	})
	router.Get("/doc/*", httpSwagger.Handler(
		httpSwagger.URL("/doc/openapi.yaml"),
	))

	const addr = ":8080"
	log.Printf("API server listening on %s", addr)

    if err := http.ListenAndServe(addr, router); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}
