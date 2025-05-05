package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/ahsan0098/gofirst/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set in .env file")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         300,
	}).Handler)

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handleReadiness)
	v1Router.Get("/error", handleErr)
	v1Router.Post("/user", apiCfg.handleCreateUser)
	v1Router.Get("/user", apiCfg.authMiddleware(apiCfg.HandleGetUser))

	v1Router.Post("/feed", apiCfg.authMiddleware(apiCfg.handleCreateFeed))
	v1Router.Get("/feed", apiCfg.HandleGetFeed)
	v1Router.Get("/user-feeds", apiCfg.authMiddleware(apiCfg.HandleGetUserFeeds))

	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	router.Mount("/api/v1", v1Router)
	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf("Starting server on port %s\n", portString)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	} else {
		fmt.Println("Server started successfully")
	}
}
