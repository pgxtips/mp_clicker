package models

import "time"

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

func create_user(email string, username string, password string){
	// validate email is not in use
	// validate username is not in use
	// create user
}

func get_user(username string){
	// return user data
}

func login_user(username string, password string){
	// validate username and password
	// generate forever token 
}
