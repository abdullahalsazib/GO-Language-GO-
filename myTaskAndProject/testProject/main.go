package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World")

	c := color.New(color.FgGreen).Add(color.Underline)

	/*
		    r := mux.NewRouter()
			r.HandleFunc("/", homeServer)

			// Create a server |> GO =>
			fmt.Println("Server running on http://localhost:8080")
			http.ListenAndServe(":8080", r)
	*/

	r := mux.NewRouter()

	staticDir := "./static"

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))

	// Starting server
	fmt.Println("Server running .....")
	fmt.Println("Server running .....")
	c.Println("http://localhost:8080")
	http.ListenAndServe(":8080", r)

}

// func homeServer(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Hello, World..!</h1>"))
// }
