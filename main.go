package main

import (
	"backend-sia/config"
	"backend-sia/databases"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	app := http.NewServeMux()
	if err := config.LoadEnv(".env"); err != nil {
		fmt.Println("Error: Failed to load file env.......")
		return
	}
	databases.ConnectDB()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"Content-Type", "application/json"},
	}).Handler

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	app.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to API Sistem Akademik")
	})
	server := &http.Server{
		Addr:    port,
		Handler: corsHandler(app),
	}

	log.Printf("Server running in port %s.......", port)
	log.Fatal(server.ListenAndServe())
}
