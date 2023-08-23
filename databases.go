package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	Login    = "root"
	Password = "root"
	DBip     = "127.0.0.1:3306"
	DBName   = "test_bd"
)

func addUser(userEmail string, userPassword string, phoneNumber string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	sqlRequest := `INSERT INTO users (Password, Email, PhoneNumber, IsAdmin) VALUES (?, ?, ?, ?)`
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", Login, Password, DBip, DBName))
	res, err := db.Query(sqlRequest, hashedPassword, userEmail, phoneNumber, false)
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()
	defer db.Close()
}