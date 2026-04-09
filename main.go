package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{}

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTodos(w, r)
		} else if r.Method == http.MethodPost {
			addTodo(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
