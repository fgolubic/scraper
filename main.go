package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/fgolubic/scraper/handlers/v1"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	dotenv "github.com/joho/godotenv"
)

func main() {

	fmt.Println("Starting server...")

	dotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the environment!")
	}

	fmt.Printf("...on port %s\n", port)
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// TODO: check what option method is
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		// TODO: check what Link header does
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,	
}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlers.HandleReadiness)
	v1Router.Get("/err", handlers.HandleError)

	router.Mount("/v1", v1Router)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
