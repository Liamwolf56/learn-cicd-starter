package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

type apiConfig struct{}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("PORT not set. Defaulting to %s", port)
	}

	router := chi.NewRouter()
	apiCfg := apiConfig{}

	router.Get("/ready", handlerReadiness)

	v1Router := chi.NewRouter()
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
	v1Router.Get("/notes", apiCfg.middlewareAuth(apiCfg.handlerNotesGet))
	v1Router.Post("/notes", apiCfg.middlewareAuth(apiCfg.handlerNotesCreate))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Serving on port: %s\n", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
