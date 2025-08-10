package main 

import (
	"os"
	"fmt"
	"log/slog"
	"server/internal"
    "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	appdata := internal.NewAppData()
	internal.SetAppData(appdata)

	releaseType := os.Getenv("RELEASE_TYPE")

	if releaseType == "" {
		slog.Error("failed to retrieve env: RELEASE_TYPE")
		os.Exit(1)
	}

	var dbPath string
	switch releaseType	{
		case "DEV":
			dbPath = "./data/local.db"
		case "PROD":
			dbPath = "./data/data.db"
		default:
			slog.Error(fmt.Sprintf("Invalid RELEASE_TYPE: %s", releaseType))
			os.Exit(1)
	}

	appdata.DBpath = &dbPath
	appdata.DB = internal.InitDB()
}
