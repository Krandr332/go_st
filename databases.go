	// databases.go
	package main

	import (
		"database/sql"
		"fmt"

		_ "github.com/lib/pq"
	)



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
	
	func heckForUserInSystem(username string) int {
		sqlRequest := `SELECT COUNT(*) FROM "public.user" WHERE username = $1;`
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/base?sslmode=disable")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
	
		var count int
		err = db.QueryRow(sqlRequest, username).Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	
		return count
	}
	
	func checkDataUser(username, password string) (bool, error) {
    sqlRequest := `SELECT password FROM "public.user" WHERE username = $1;`
    db, err := sql.Open("postgres", "postgres://postgres:1@localhost/base?sslmode=disable")
    if err != nil {
        return false, err
    }
    defer db.Close()

    var passDB string
    err = db.QueryRow(sqlRequest, username).Scan(&passDB)
    if err != nil {
        if err == sql.ErrNoRows {
            // Пользователь с указанным именем не найден.
            return false, nil
        }
        return false, err
    }

    if password == passDB {
        return true, nil
    }

    return false, nil
}