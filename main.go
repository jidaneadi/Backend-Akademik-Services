package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-sia/configs"
	"project-sia/controllers/authcontrollers"
	"project-sia/databases"
	"project-sia/exceptions"
	"project-sia/middlewares"
	"project-sia/repository/siswarepository"
	"project-sia/routes"
	"project-sia/services/authservices"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

func errorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				exceptions.ErrorHandler(w, r, err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func main() {
	app := http.NewServeMux()
	if err := configs.LoadEnv(".env"); err != nil {
		fmt.Println("Error: Failed to load file env.......")
		return
	}
	databases.ConnectDB()
	validate := validator.New()

	siswarepository := siswarepository.NewSiswaRepository()
	authService := authservices.NewAuthServices(validate, siswarepository)
	authController := authcontrollers.NewAuthControllers(authService)
	routes.AuthRoutes(authController, app)

	handler := middlewares.LoggingMiddleware(app)
	handler = errorHandlerMiddleware(handler)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"Content-Type", "application/json"},
	}).Handler(handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	app.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to API Sistem Akademik")
	})
	server := &http.Server{
		Addr:    port,
		Handler: corsHandler,
	}

	log.Printf("Server running in port %s.......", port)
	log.Fatal(server.ListenAndServe())
}
