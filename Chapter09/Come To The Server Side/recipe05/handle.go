package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "User GET")
		}
		if r.Method == http.MethodPost {
			fmt.Fprintln(w, "User POST")
		}
	})

	// separate handler
	itemMux := http.NewServeMux()
	itemMux.HandleFunc("/items/clothes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Clothes")
	})
	mux.Handle("/items/", itemMux)

	// Admin handlers
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/ports", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ports")
	})

	mux.Handle("/admin/",
		http.StripPrefix("/admin", adminMux))

	// Default server
	http.ListenAndServe(":8080", mux)

}
