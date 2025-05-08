package main

import (
	"fmt"

	"github.com/maxBRT/internal/database"
)

type apiConfig struct {
	database *database.Queries
}

func main() {
	fmt.Println("hello world")
}
