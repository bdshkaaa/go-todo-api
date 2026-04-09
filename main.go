package main

import (
	"log"
	"net/http"

	"go-todo-api/handlers"
)

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetTodos(w, r)
		} else if r.Method == http.MethodPost {
			handlers.AddTodo(w, r)
		}
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
