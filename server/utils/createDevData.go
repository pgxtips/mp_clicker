package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"server/internal"
	"server/internal/models"
)

func main(){
	appdata := internal.NewAppData()
	internal.SetAppData(appdata)

	rootPath := internal.GetRootDir()
	dbPath := filepath.Join(rootPath, "data", "local.db")

	// wipe old data
	if err := os.Remove(dbPath); err != nil && !errors.Is(err, os.ErrNotExist){
		slog.Error(fmt.Sprintf("Error removing local db file: %v", err))
		os.Exit(1)
	}

	appdata.DB = internal.InitDB()

	// create user data
	u1, err1 := models.CreateUser(appdata.DB, "dev1@dev.com", "dev1", "password1")
	if err1 != nil{
		fmt.Printf("There was an issue: %v", err1)
		os.Exit(1)
	}
	fmt.Println("user 1:", *u1)

	u2, err2 := models.CreateUser(appdata.DB, "dev2@dev.com", "dev2", "password2")
	if err2 != nil{
		fmt.Printf("There was an issue: %v", err2)
		os.Exit(1)
	}
	fmt.Println("user 2:", *u2)

	u3, err3 := models.CreateUser(appdata.DB, "dev3@dev.com", "dev3", "password3")
	if err3 != nil{
		fmt.Printf("There was an issue: %v", err3)
		os.Exit(1)
	}
	fmt.Println("user 3:", *u3)

	// create world data
	w1, err1 := models.CreateWorld(appdata.DB, u1.Id, "dev1 world 1")
	if err1 != nil{
		fmt.Printf("There was an issue: %v", err1)
		os.Exit(1)
	}
	fmt.Println("world 1:", *w1)

	w2, err2 := models.CreateWorld(appdata.DB, u2.Id, "dev2 world 1")
	if err1 != nil{
		fmt.Printf("There was an issue: %v", err1)
		os.Exit(1)
	}
	fmt.Println("world 2:", *w2)

	w3, err3 := models.CreateWorld(appdata.DB, u3.Id, "dev3 world 1")
	if err1 != nil{
		fmt.Printf("There was an issue: %v", err1)
		os.Exit(1)
	}
	fmt.Println("world 3:", *w3)
}
