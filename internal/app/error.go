package app

import (
	"encoding/json"
	"net/http"
)

type errorMessage struct {
	Message string `json:"message"`
}

func (app *App) handleError(w http.ResponseWriter, code int, msg string) {
	e := errorMessage{msg}
	jsonErr, _ := json.Marshal(e)

	w.WriteHeader(code)
	w.Write(jsonErr)
}
