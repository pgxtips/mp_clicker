package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	appdata := GetAppData()
	dbPath := *appdata.DBpath

	// ensure data dir exists
    os.MkdirAll("./data/", 0755)

	// open/create file
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} 

	InitSchema(db)

	return db
}

func InitSchema(db *sql.DB){
	// create users table
	// dw post alpha, users pwd's will be salted and hashed (chillll)
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS 'users' 
		(
			'id' INTEGER PRIMARY KEY AUTOINCREMENT, 
			'email' VARCHAR(255) NOT NULL,
			'username' TEXT NOT NULL,
			'password' TEXT NOT NULL,
			'created_at' DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			'deleted' INTEGER NOT NULL DEFAULT 0 
		)
	`)
    if err != nil {
		fmt.Printf("failed to create 'users' table: %v", err)
        os.Exit(1)
    }

	// create gamedata table 
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS 'worlddata' 
		(
			'id' INTEGER PRIMARY KEY AUTOINCREMENT, 
			'name' TEXT NOT NULL,
			'owner' FORIEGN KEY NOT NULL,
			'world_data' TEXT,
			'modifier_data' TEXT
			'created_at' DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			'deleted' INTEGER NOT NULL DEFAULT 0 
		)
	`)
    if err != nil {
		fmt.Printf("failed to create 'worlddata' table: %v", err)
        os.Exit(1)
	}

	// create users_worlds resolver table 
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS 'users_worlddata' 
		(
			'id' INTEGER PRIMARY KEY AUTOINCREMENT, 
			'uid' INTEGER FORIEGN KEY, 
			'wid' INTEGER FORIEGN KEY
		)
	`)
    if err != nil {
		fmt.Printf("failed to create 'users_worlds' table: %v", err)
        os.Exit(1)
	}
}
