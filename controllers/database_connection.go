package controllers

import (
	"database/sql"
	"github.com/dastardlyjockey/hngx-2/internal/database"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func DBInstance() *database.Queries {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment variable")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open database connection, error: ", err)
	}

	db := database.New(connection)

	return db
}

var ApiCfg = &ApiConfig{
	DB: DBInstance(),
}
