package main 

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"server/internal"
	"server/internal/models"
)

func main(){
	appdata := internal.NewAppData()
	internal.SetAppData(appdata)

	var dbPath = "./data/local.db"
	appdata.DBpath = &dbPath

	// wipe old data
	if err := os.Remove("./data/local.db"); err != nil && !errors.Is(err, os.ErrNotExist){
		slog.Error(fmt.Sprintf("Error removing local db file: %v", err))
		os.Exit(1)
	}

	appdata.DB = internal.InitDB()

	// insert user data
	models.CreateUser(appdata.DB, "dev1@dev.com", "dev1", "password1")
	models.CreateUser(appdata.DB, "dev2@dev.com", "dev2", "password2")
	models.CreateUser(appdata.DB, "dev3@dev.com", "dev3", "password3")
}
