package main

import (
	"fmt"
	"net/http"
)

func serveLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Serve the login page
	http.ServeFile(w, r, "templates/login.html")
}

func loginAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Basic authentication logic (replace with your own authentication logic)
	if username == "user" && password == "password" {
		fmt.Fprintln(w, "Login successful")
	} else {
		http.Error(w, "Login failed", http.StatusUnauthorized)
	}
}
