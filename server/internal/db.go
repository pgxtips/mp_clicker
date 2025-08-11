package internal

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
    "github.com/jmoiron/sqlx"
)

func InitDB() *sqlx.DB {
	appdata := GetAppData()
	dbPath := *appdata.DBpath

	// ensure data dir exists
    os.MkdirAll("./data/", 0755)

	// open/create file
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} 

	InitSchema(db)

	return db
}

func InitSchema(db *sqlx.DB){
	// create users table
	// dw post alpha, users pwd's will be salted and hashed (chillll)
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users 
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			email TEXT NOT NULL,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted INTEGER NOT NULL DEFAULT 0 
		)
	`)
    if err != nil {
		fmt.Printf("failed to create 'users' table: %v", err)
        os.Exit(1)
    }

	// create gamedata table 
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS worlddata 
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			name TEXT NOT NULL,
			owner_user_id INTEGER NOT NULL REFERENCES users(id),
			modifier_data TEXT DEFAULT '{}',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted INTEGER NOT NULL DEFAULT 0 
		)
	`)
    if err != nil {
		fmt.Printf("failed to create 'worlddata' table: %v", err)
        os.Exit(1)
	}

	// create users_worlds resolver table 
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users_worlddata 
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			uid INTEGER NOT NULL REFERENCES users(id), 
			wid INTEGER NOT NULL REFERENCES worlddata(id) 
		)
	`)
    if err != nil {
		fmt.Printf("failed to create 'users_worlds' table: %v", err)
        os.Exit(1)
	}
}
