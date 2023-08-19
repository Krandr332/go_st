package main

import (
	"fmt"
	"net/http"
)



func main() {
	fmt.Println("http://localhost:8080")
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexPage)
	http.ListenAndServe(":8080", nil)
}
