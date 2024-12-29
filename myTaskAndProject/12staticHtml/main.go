package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	/*
		fmt.Println("Welcome This is Static Html - Server")
		fs := http.FileServer(http.Dir("./static"))
		http.Handle("/", fs)

		// Start Server
		fmt.Println("Server open on http://localhost:8080")
		http.ListenAndServe(":8080", fs)
	*/

	/*
		r := mux.NewRouter()
		// Serve static files from the "static" directory

		fs := http.FileServer(http.Dir("./static"))
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

		// define API route
		r.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		})

		// starting server
		fmt.Println("Server running on http://location:8080")
		http.ListenAndServe(":8080", r)
	*/

	// Define a route to serve dynamic HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("static/index.html"))
		temp.Execute(w, nil)
	})

	// Serve static assets html, css, js etc.....
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// start the server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
