package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Using Gorilla Mux")
	r := mux.NewRouter()

	r.HandleFunc("/", homeServe)
	r.HandleFunc("/about", aboutServe)
	r.HandleFunc("/user/{id}", userHandler) // dynamic route

	// starting server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func homeServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style='text-align: center;'>hello this is home page</h1><a href='/about'>about</a>"))
	fmt.Fprintln(w, "Hello home page... !")

}

func aboutServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style='text-align: center;'>hello this is about page</h1><a href='/'>about</a>"))
	fmt.Fprintln(w, "Hello home page... !")

}

// user handler with dynamic route
func userHandler(w http.ResponseWriter, r *http.Request) {
	perams := mux.Vars(r)
	id := perams["id"]

	fmt.Fprintf(w, "User id %s\n", id)
}
