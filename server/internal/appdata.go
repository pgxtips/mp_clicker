package internal

import (
	"database/sql"
)

type AppData struct {
	ReleaseType *string 
	DBpath *string 
	DB *sql.DB
}

// initialised on startup (elsewhere)
var globalAppData *AppData

func NewAppData() *AppData {
	globalAppData = &AppData{
		ReleaseType: nil,
		DBpath: nil,
		DB: nil,
	}

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
