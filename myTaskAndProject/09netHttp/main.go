package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello here is my net http routes")

	// define routes
	http.HandleFunc("/", homeServe)
	http.HandleFunc("/about", aboutServe)

	// Start a server
	fmt.Println("server runnling on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Go home page</h1> <a href='/about'>about</a> "))
	fmt.Fprintln(w, "Welcome to the home page!")
}
func aboutServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Go about page</h1><a href='/'>home</a>"))
	fmt.Fprintln(w, "Welcome to the about page!")
}
