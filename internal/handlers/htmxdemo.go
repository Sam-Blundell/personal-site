package handlers

import (
	"html/template"
	"net/http"
)

func HtmxDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/htmxdemo.html",
	)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Heading string
	}{
		Title:   "HTMX Demo",
		Heading: "HTMX + Go Templates",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error serving template", http.StatusInternalServerError)
	}
}
