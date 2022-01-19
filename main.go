package main

import (
	"log"

	"github.com/MrTj458/fiber-api-todo/http"
	"github.com/MrTj458/fiber-api-todo/inmemory"
)

// @title Todo api
// @version 1.0
func main() {
	s := http.NewServer(3000)

	todoService := inmemory.NewTodoService()

	s.TodoService = todoService

	log.Fatal(s.Run())
}
