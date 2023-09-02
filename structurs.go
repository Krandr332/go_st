	// structurs.go

package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	UserID   int
	Username string
	Email    string
	Password string
}

type Friend struct {
	FriendshipID int
	UserID1      int
	UserID2      int
	Status       string
}

type Post struct {
	PostID     int
	UserID     int
	PhotoURL   string
	Text       string
	CreatedAt  time.Time
}

type Message struct {
	MessageID   int
	SenderID    int
	RecipientID int
	Content     string
	SentAt      time.Time
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
