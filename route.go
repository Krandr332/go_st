// route.go
package main

import (
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
	t.Execute(w, nil)

}
func registerPage(w http.ResponseWriter, r *http.Request) {
	// data:=""
	err:= r.ParseForm()
	if err!=nil{
		log.Fatal(err)
	}
	if r.Method == "POST"{
		username:=r.FormValue("username")
		password1 :=r.FormValue("password1")
		password2 :=r.FormValue("password2")
		if password1 == password2{
			CreateUserAccaunt(username,password1)
			if err !=nil{
				log.Println(err)
				http.Error(w,"qq",http.StatusInternalServerError)
			}
			
		}
	}
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