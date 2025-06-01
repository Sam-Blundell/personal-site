package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

func htmxdemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/htmxdemo.html",
	)
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
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
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		return
	}
}

func gettime(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/partials/time.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/htmxdemo", htmxdemo)
	mux.HandleFunc("/time", gettime)

	fmt.Println("Starting server on http://localhost:8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
