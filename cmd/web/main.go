package main

import (
	"fmt"
	"net/http"

	"github.com/Sam-Blundell/personal-site/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/htmxdemo", handlers.HtmxDemo)
	mux.HandleFunc("/time", handlers.GetTime)

	fmt.Println("Starting server on http://localhost:8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
