package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World")
	/*
		    r := mux.NewRouter()
			r.HandleFunc("/", homeServer)

			// Create a server |> GO =>
			fmt.Println("Server running on http://localhost:8080")
			http.ListenAndServe(":8080", r)
	*/

	r := mux.NewRouter()

	staticDir := "/static"

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))

	// Starting server
	fmt.Println("Server running ..... http://localhost:8080")
	http.ListenAndServe(":8080", r)

}

// func homeServer(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Hello, World..!</h1>"))
// }
