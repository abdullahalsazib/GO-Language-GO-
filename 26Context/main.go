package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello, Context")

	http.HandleFunc("/", Hello) // Fix: Use correct function name
	log.Fatal(http.ListenAndServe(":8080", nil)) // Log fatal only for server startup errors
}

func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("Server: Hello handle started..")
	defer fmt.Println("Server: Hello handle end..")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			log.Println("Error:", err) // Fix: Use log.Println instead of log.Fatal
		}
		internalError := http.StatusInternalServerError // Fix: Correct variable name
		http.Error(w, err.Error(), internalError)
	}
}
