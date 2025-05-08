package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/maxBRT/assets-man/internal/database"
)

type apiConfig struct {
	database *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)

	}

	dbQueries := database.New(db)

	serveMux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	config := &apiConfig{
		database: dbQueries,
	}

	fileServer := http.FileServer(http.Dir("."))

	serveMux.Handle("/", fileServer)

	serveMux.HandleFunc("POST /api/users", config.addUser)

	server.ListenAndServe()
}
