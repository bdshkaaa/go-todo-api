package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-todo-api/models"
	"go-todo-api/storage"
)

var todos []models.Todo
var nextID = 1

func init() {
	loaded, _ := storage.LoadTodos()
	todos = loaded

	for _, t := range todos {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo

	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newTodo.ID = nextID
	nextID++

	todos = append(todos, newTodo)
	storage.SaveTodos(todos)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			storage.SaveTodos(todos)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(w, r)
		case http.MethodPost:
			addTodo(w, r)
		case http.MethodDelete:
			deleteTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
