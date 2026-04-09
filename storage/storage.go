package storage

import (
	"encoding/json"
	"os"

	"go-todo-api/models"
)

var FilePath = "data/todos.json"

func LoadTodos() ([]models.Todo, error) {
	file, err := os.ReadFile(FilePath)
	if err != nil {
		return []models.Todo{}, nil // если файла нет — ок
	}

	var todos []models.Todo
	json.Unmarshal(file, &todos)

	return todos, nil
}

func SaveTodos(todos []models.Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(FilePath, data, 0644)
}
