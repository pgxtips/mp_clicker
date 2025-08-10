package models

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct{
	// db schema
	Id int 
	Email string
	Username string
	Created_at time.Time 
	Deleted bool 

	// internal use only
	IsAuthed bool 
}

func CreateUser(db *sql.DB, email string, username string, password string){
	// validate email is not in use
	// validate username is not in use
	// create user
	_, err := db.Exec(`
		INSERT INTO users (email, username, password)
		VALUES (?, ?, ?)
	`, email, username, password)
	if err != nil {
		fmt.Printf("failed to insert user: %v\n", err)
	}
}

func GetUser(username string){
	// return user data
}

func LoginUser(username string, password string){
	// validate username and password
	// generate forever token 
}
