package main 

import (
	"os"
	"log/slog"
	"server/internal"

    "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	releaseType := os.Getenv("RELEASE_TYPE")

	if releaseType != "" {
		slog.Error("failed to retrieve env: RELEASE_TYPE")
		os.Exit(1)
	}

	internal.CreateDB()
}
