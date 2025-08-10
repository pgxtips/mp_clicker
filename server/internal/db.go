package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	_, err := os.Stat("./data/data.db")
	if errors.Is(err, os.ErrNotExist) {
		CreateDB()
	}
}

func initTables(db *sql.DB){
	// create users table
	// dw post alpha, users pwd's will be salted and hashed (chillll)
	_, err := db.Exec(`
		CREATE TABLE 'users' 
		(
			'id' INTEGER PRIMARY KEY AUTOINCREMENT, 
			'email' VARCHAR(255) NOT NULL,
			'username' VARCHAR(255) NULL,
			'password' VARCHAR(255) NULL,
			'created_at' DATETIME NULL,
			'deleted' BOOL FALSE
		)
	`)
    if err != nil {
		fmt.Printf("failed to create 'users' table: %v", err)
        os.Exit(1)
    }

	// create gamedata table 
	_, err = db.Exec(`
		CREATE TABLE 'worlddata' 
		(
			'id' INTEGER PRIMARY KEY AUTOINCREMENT, 
			'name' TEXT NOT NULL,
			'owner' FORIEGN KEY NOT NULL,
			'created_at' TEXT NOT NULL,
			'world_data' TEXT NOT NULL,
			'modifier_data' TEXT NOT NULL
		)
	`)
    if err != nil {
		fmt.Printf("failed to create 'worlddata' table: %v", err)
        os.Exit(1)
	}

	// create users_worlds resolver table 
	_, err = db.Exec(`
		CREATE TABLE 'users_worlddata' 
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

func CreateDB() {
    os.MkdirAll("./data/", 0755)
    os.Create("./data/data.db")

    db, err := sql.Open("sqlite3", "./data/data.db")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }; 
	initTables(db)
	db.Close()
}
