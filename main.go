package main

import (
	"fmt"
	"net/http"
)



func main() {
	fmt.Println("http://localhost:8080")
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/register",registerPage)
	http.HandleFunc("/home",homePage)

	http.ListenAndServe(":8080", nil)
}
