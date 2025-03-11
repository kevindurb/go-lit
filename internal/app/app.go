package app

import (
	"kevindurb/go-lit/internal/entities"
	"net/http"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (app *App) Run() {
	app.DB.AutoMigrate(&entities.Todo{})

	http.HandleFunc("POST /todos", app.handleCreateTodo)
	http.HandleFunc("GET /todos", app.handleListTodos)
	http.ListenAndServe(":3000", nil)
}
