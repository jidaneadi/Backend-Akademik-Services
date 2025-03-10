package main

import (
	"backend-sia/config"
	"backend-sia/database"
	"backend-sia/services"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func homeRootControllers(c http.ResponseWriter, r *http.Request) error {
	return services.MessageSucces(c, "Welcome to api sistem informasi siswa", false)
}

func withErrorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	database.ConnectDB()
	app := http.NewServeMux()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET, POST, PUT, DELETE"},
	}).Handler
	app.HandleFunc("/", withErrorHandler(homeRootControllers))
	port := config.Renderenv("PORT")
	if port == "" {
		port = "3000"
	}
	server := &http.Server{
		Addr:    port,
		Handler: corsHandler(app),
	}

	log.Printf("Server running in port %s ........", port)
	log.Fatal(server.ListenAndServe())
}
