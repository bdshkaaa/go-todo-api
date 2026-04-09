package handlers

import (
	"encoding/json"
	"net/http"

	"go-todo-api/models"
	"go-todo-api/storage"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, _ := storage.LoadTodos()
	json.NewEncoder(w).Encode(todos)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo

	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todos, _ := storage.LoadTodos()
	newTodo.ID = len(todos) + 1

	todos = append(todos, newTodo)
	storage.SaveTodos(todos)

	json.NewEncoder(w).Encode(newTodo)
}
