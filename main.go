package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	tasks = []Task{
		{ID: 1, Name: "Learn Go"},
		{ID: 2, Name: "Build REST API"},
		{ID: 3, Name: "Learn Docker"},
	}

	mu sync.Mutex
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status":  "UP",
		"message": "Go Task API is running",
	}

	json.NewEncoder(w).Encode(response)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		getTasks(w, r)

	case http.MethodPost:
		createTask(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {

	idString := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		getTask(w, r, id)

	case http.MethodPut:
		updateTask(w, r, id)

	case http.MethodDelete:
		deleteTask(w, r, id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request, id int) {

	mu.Lock()
	defer mu.Unlock()

	for _, task := range tasks {

		if task.ID == id {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(task)

			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func createTask(w http.ResponseWriter, r *http.Request) {

	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request, id int) {

	var updatedTask Task

	err := json.NewDecoder(r.Body).Decode(&updatedTask)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i := range tasks {

		if tasks[i].ID == id {

			tasks[i].Name = updatedTask.Name

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(tasks[i])

			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, r *http.Request, id int) {

	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {

		if task.ID == id {

			tasks = append(tasks[:i], tasks[i+1:]...)

			w.WriteHeader(http.StatusNoContent)

			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/tasks", tasksHandler)

	http.HandleFunc("/tasks/", taskHandler)

	log.Println("Go Task API running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
