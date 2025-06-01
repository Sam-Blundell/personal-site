package handlers

import (
	"html/template"
	"net/http"
	"time"
)

func GetTime(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/partials/time.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Time string
	}{
		Time: time.Now().Format("15:04:05 MST"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering partial", http.StatusInternalServerError)
	}
}
