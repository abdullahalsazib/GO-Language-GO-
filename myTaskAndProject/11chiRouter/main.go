package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("Welcome in here Go - Chi router")
	r := chi.NewRouter()

	// add middlelware
	r.Use(middleware.Logger)

	r.Get("/", homeServe)
	r.Get("/about", aboutServe)
	r.Get("/blog", blogServe)
	r.Get("/user/{id}", users)

	// Server starting
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func homeServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello home</h1>"))
}
func aboutServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello about</h1>"))
}
func blogServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello blog</h1>"))
}
func users(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("User id is %s\n", id)))
}
