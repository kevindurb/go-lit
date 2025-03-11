package app

import (
	"encoding/json"
	"kevindurb/go-lit/internal/entities"
	"net/http"
)

type createTodoBody struct {
	Description string `json:"description"`
}

func (app *App) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var b createTodoBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		app.handleError(w, http.StatusBadRequest, "Bad request")
		return
	}

	t := entities.Todo{Description: b.Description}
	result := app.DB.Create(&t)
	if result.Error != nil {
		app.handleError(w, http.StatusInternalServerError, "Error creating todo")
		return
	}

	response, _ := json.Marshal(t)
	w.Write(response)
}

func (app *App) handleListTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var t []entities.Todo
	result := app.DB.Find(&t)
	if result.Error != nil {
		app.handleError(w, http.StatusInternalServerError, "Error getting todos")
	}

	response, _ := json.Marshal(t)
	w.Write(response)
}
