package main

import (
	"database/sql"
	"embed"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/bootdotdev/learn-cicd-starter/internal/database"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type apiConfig struct {
	DB *database.Queries
}

//go:embed static/*
var staticFiles embed.FS

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: .env not loaded: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not set, defaulting to 8080")
		port = "8080"
	}

	apiCfg := apiConfig{}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		db, err := sql.Open("libsql", dbURL)
		if err != nil {
			log.Fatal(err)
		}
		apiCfg.DB = database.New(db)
		log.Println("Connected to database!")
	} else {
		log.Println("No DATABASE_URL; running in non-DB mode")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := staticFiles.Open("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		if _, err = io.Copy(w, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	v1 := chi.NewRouter()
	if apiCfg.DB != nil {
		v1.Post("/users", apiCfg.handlerUsersCreate)
		v1.Get("/users", apiCfg.middlewareAuth(authedHandler(apiCfg.handlerUsersGet)))
		v1.Get("/notes", apiCfg.middlewareAuth(authedHandler(apiCfg.handlerNotesGet)))
		v1.Post("/notes", apiCfg.middlewareAuth(authedHandler(apiCfg.handlerNotesCreate)))
	}
	v1.Get("/healthz", handlerReadiness)
	router.Mount("/v1", v1)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Serving on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
