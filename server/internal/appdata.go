package internal

import (
	"os"
	"path/filepath"
    "github.com/jmoiron/sqlx"
)

type AppData struct {
	DBpath *string 
	DB *sqlx.DB
}

func GetRootDir() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// walk up until we find go.mod (repo root)
	for {
		if _, err := os.Stat(filepath.Join(root, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(root)
		if parent == root {
			panic("could not find repo root (go.mod)")
		}
		root = parent
	}

	return root
}

// initialised on startup (elsewhere)
var globalAppData *AppData

func NewAppData() *AppData {
	rootPath := GetRootDir()
	dbPath := filepath.Join(rootPath, "data", "local.db")
	globalAppData = &AppData{
		DBpath: &dbPath,
		DB: nil,
	}

	globalAppData.DB = InitDB()
	return globalAppData
}

func SetAppData(a *AppData) {
	globalAppData = a
}

func GetAppData() *AppData {
    if globalAppData == nil {
        panic("GlobalAppData not initialized")
    }
    return globalAppData
}
