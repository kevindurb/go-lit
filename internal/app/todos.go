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
	json.NewDecoder(r.Body).Decode(&b)
	t := entities.Todo{}
	app.DB.Create(&t)
	response, _ := json.Marshal(t)
	w.Write(response)
}
