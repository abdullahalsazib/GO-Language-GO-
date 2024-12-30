package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

// Create a new color object
var c1 = color.New(color.FgCyan).Add(color.Underline)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks []Task
var nextId int = 1

func main() {
	fmt.Println("That is the API Created by Jack Sparrow...")
	r := mux.NewRouter()

	// Server Static file
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Routers
	r.HandleFunc("/tasks", getTasks).Methods("GET")

}

// Handler's

// All Task's
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// get a task by id
func getTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

}

// create new Task
func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var task
}

func checkError(err error) {
	if err != nil {
		c1.Printf("Somthing is Wrong %v", err)
	}
}
