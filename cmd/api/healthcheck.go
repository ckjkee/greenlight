package main

import (
	"net/http"
)

// Можно использовать сырой текст для отправки HTTP ответов, не лучшая практика, но где то можно применять
// js := `{"status: "available", "environment": %q, "version": %q}`
// 	js = fmt.Sprintf(js, app.cfg.env, version)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(js))

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.cfg.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
