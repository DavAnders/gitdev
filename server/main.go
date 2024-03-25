package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/SamiZeinsAI/gitdev/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	port         string
	DB           *database.Queries
	jwtSecret    string
	clientID     string
	clientSecret string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbUrl := os.Getenv("POSTGRES")
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	apiCfg := apiConfig{
		port:         os.Getenv("PORT"),
		DB:           database.New(conn),
		jwtSecret:    os.Getenv("JWT_SECRET"),
		clientID:     os.Getenv("GITHUB_CLIENT_ID"),
		clientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
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
	v1Router.Get("/err", apiCfg.handlerErr)
	v1Router.Get("/auth/{provider}/callback", apiCfg.handlerGitHubCallback)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: router,
	}
	log.Printf("Serving on port: %s\n", apiCfg.port)
	log.Fatal(srv.ListenAndServe())

}
