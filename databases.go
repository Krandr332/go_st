	// databases.go
	package main

	import (
		"database/sql"
		"fmt"
		// "log"

		_ "github.com/lib/pq"
	)



	// var db *sql.DB

	// const (
	// 	host = "localhost"
	// 	port = 5432
	// 	user = "postgres"
	// 	password = "1"
	// 	dbname = "base"
	// )
	func CreateUserAccaunt(username,email,password string) error {
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/base?sslmode=disable")
		if err != nil {
			return err
		}
		defer db.Close()
	
		// Проверяем соединение с базой данных
		err = db.Ping()
		if err != nil {
			return err
		}
	
		// Создаем нового пользователя в таблице users
		_, err = db.Exec(`INSERT INTO "public.user" (username, email, password) VALUES ($1, $2, $3)`, username, email, password)
		if err != nil {
			return err
		}
	
		fmt.Println("Пользователь успешно создан")
	
		return nil
	}
	
	