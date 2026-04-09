# 🟦 Go Todo API

Simple REST API for managing tasks, built with Go.

---

## 🚀 Features

* Create tasks
* Get all tasks
* Persistent storage (data saved to JSON file)
* Clean project structure (handlers, models, storage)

---

## 📁 Project Structure

```
.
├── main.go
├── models/
│   └── todo.go
├── handlers/
│   └── todo.go
├── storage/
│   └── storage.go
├── data/
│   └── todos.json
└── go.mod
```

---

## ⚙️ Installation

```bash
git clone https://github.com/bdshkaaa/go-todo-api.git
cd go-todo-api
go mod tidy
```

---

## ▶️ Run server

```bash
go run main.go
```

Server will start at:

http://localhost:8080

---

## 📡 API Endpoints

### GET /todos

Returns all tasks.

---

### POST /todos

Creates a new task.

Example request body:

```json
{
  "title": "Learn Go"
}
```

---

### DELETE /todos?id=1

Deletes a task by ID.

---

## 💾 Data Storage

All tasks are stored in:

```
data/todos.json
```

---

## 🛠 Technologies

* Go (net/http)
* JSON file storage

---

## 📌 Author

GitHub: https://github.com/bdshkaaa
