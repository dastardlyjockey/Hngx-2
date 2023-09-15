package main

import (
	"github.com/dastardlyjockey/hngx-2/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not found in the environment variable")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
	}))

	v1Router := chi.NewRouter()

	//Route
	routes.UserRoute(v1Router)

	//mounting to api path
	router.Mount("/api", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
