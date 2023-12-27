package main

import (
	"context"
	"net/http"
	"time"

	"github.com/hugogarcia/microservices/logger-service/data"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload jsonPayload
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*15)
	defer cancel()

	_ = app.readJSON(w, r, &requestPayload)

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(ctx, event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "logged",
	})

}
