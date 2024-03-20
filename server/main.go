package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	port string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	apiCfg := apiConfig{
		port: os.Getenv("PORT"),
	}

	if apiCfg.port == "" {
		apiCfg.port = "80"
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()

	v1Router.Get("/readiness", apiCfg.handlerReadiness)
	srv := &http.Server{
		Addr:    apiCfg.port,
		Handler: mux,
	}
	log.Printf("Serving on port: %s\n", apiCfg.port)
	log.Fatal(srv.ListenAndServe())

}
