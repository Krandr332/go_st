package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	UserID   int
	Username string
	Email    string
	Password string
}

type Database struct {
	*sql.DB
}

func NewDatabase(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (db *Database) RegisterUser(username, email, password string) (int, error) {
	var userID int
	err := db.QueryRow("INSERT INTO public.user (username, email, password) VALUES ($1, $2, $3) RETURNING user_id",
		username, email, password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func main() {
	connStr := "host=your-host dbname=your-dbname user=your-user password=your-password sslmode=disable"
	db, err := NewDatabase(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	username := "newuser"
	email := "newuser@example.com"
	password := "securepassword"

	userID, err := db.RegisterUser(username, email, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User registered with ID: %d\n", userID)
}
