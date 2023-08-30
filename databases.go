	// databases.go
	package main

	import (
		"database/sql"
		// "fmt"
		// "log"

		_ "github.com/lib/pq"
	)



	var db *sql.DB

	const (
		host = "localhost"
		port = 5432
		user = "postgres"
		password = "1"
		dbname = "base"
	)
	func CreateUserAccaunt(username,password string)  {
		
		query := "INSERT INTO public.user(username, password) VALUES ($1, $2)"
    _, err := db.Exec(query, username, password)
    CheckError(err)
	}

	func CheckError(err error){
		if err!=nil{
			panic(err)
		}
	}