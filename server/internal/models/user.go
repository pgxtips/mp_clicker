package models

import (
	"fmt"
	"time"
    "github.com/jmoiron/sqlx"
)

type User struct{
	// db schema
	Id int `db:"id"`
	Email string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
	Created_at time.Time `db:"created_at"`
	Deleted int `db:"deleted"`

	// internal use only
	IsAuthed bool 
}

func CreateUser(db *sqlx.DB, email string, username string, password string) (*User, error) {
	// validate email is not in use
	// validate username is not in use
	// create user
	u := &User{}
	err := db.Get(u, `
		INSERT INTO users (email, username, password)
		VALUES (?, ?, ?)
		RETURNING *;
	`, email, username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v\n", err)
	}

	return u, nil
}

func GetUser(db *sqlx.DB, username string) (*User, error){
	u := &User{}

	err:= db.Unsafe().Get(u, `
		SELECT * FROM users
		WHERE username = ?
		AND deleted == 0
		LIMIT 1;
	`, username)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return u, nil
}

func LoginUser(db *sqlx.DB, username string, password string) (*string, error){
	u := &User{}
	err:= db.Unsafe().Get(u, `
		SELECT * FROM users
		WHERE username = ?
		AND deleted == 0
		LIMIT 1;
	`, username)

	if err != nil {
		return nil, fmt.Errorf("user does not exist")
	}

	// validate username and password
	if u.Username == username && u.Password == password {
		// return token 
		return nil, nil
	}

	return nil, fmt.Errorf("incorrect password")
}
