package main

import (
	"html/template"
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