package main

import (
    "fmt"
    "net/http"
    "html/template"
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

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    fmt.Println("Starting server on http://localhost:8080...")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        panic(err)
    }
}

