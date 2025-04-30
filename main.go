package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         300,
	}).Handler)


	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/ready", handleReadiness)
	v1Router.HandleFunc("/error", handleErr)

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
