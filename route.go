// route.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)

}
func loginPage(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/login.html")
    if err != nil {
        panic(err)
    }

    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        isAuthenticated, err := checkDataUser(username, password)
        if err != nil {	
            http.Error(w, "Ошибка при аутентификации", http.StatusInternalServerError)
			fmt.Println("слоупок")
            return
        }

        if isAuthenticated {

			fmt.Println("все ок")
            http.Redirect(w, r, "/home", http.StatusSeeOther)
            return
        }
        errMsg := "Неправильное имя пользователя или пароль"
        t.Execute(w, errMsg)
        return
    }

    t.Execute(w, nil)
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	// data:=""
	dubleUser:=false
	err:= r.ParseForm()
	if err!=nil{
		log.Fatal(err)
	}
	if r.Method == "POST"{
		username:=r.FormValue("username")
		password1 :=r.FormValue("password1")
		password2 :=r.FormValue("password2")
		if heckForUserInSystem(username) >= 1 {
			data := "Пользователь уже есть в системе"
			fmt.Println(data)
			dubleUser = true
		} else if password1 != password2{
			dubleUser = true
			fmt.Println("пароль не тот ")
			}
		if dubleUser == false{
			err=CreateUserAccaunt(username,"nil",password1)
			if err !=nil{
				log.Println(err)
				http.Error(w,"ОШИБКА БЛЯ",http.StatusInternalServerError)
		}
		}}
	t, err := template.ParseFiles("templates/register.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}


func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("home").ParseFiles("templates/home.html", "templates/header.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
