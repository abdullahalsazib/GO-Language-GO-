package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var (
	tasks  = []Task{}
	mu     sync.Mutex
	nextID = 1
)

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		json.NewEncoder(w).Encode(tasks)
		mu.Unlock()
	case http.MethodPost:
		var task Task
		json.NewDecoder(r.Body).Decode(&task)
		mu.Lock()
		task.ID = nextID
		nextID++
		tasks = append(tasks, task)
		mu.Unlock()
		json.NewEncoder(w).Encode(task)
	case http.MethodPut:
		var updatedTask Task
		json.NewDecoder(r.Body).Decode(&updatedTask)
		mu.Lock()
		for i, t := range tasks {
			if t.ID == updatedTask.ID {
				tasks[i] = updatedTask
				break
			}
		}
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		mu.Lock()
		for i, t := range tasks {
			if t.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	}
}
