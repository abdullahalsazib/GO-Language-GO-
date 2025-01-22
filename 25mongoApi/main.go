package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("Welcome in --> Mongo-Api")
	r := router.Router()

	fmt.Println("Server is getting ready...")

	fmt.Println("Server is up and running... please visit http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
