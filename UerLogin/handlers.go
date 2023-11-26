package main

import (
	"database/sql"
	"fmt"
	"log"
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

	// Query the database for the user's credentials
	var storedPassword string
	// err = db.QueryRow("SELECT Password FROM UserLogin WHERE Username = ?", username).Scan(&storedPassword)
	err = db.QueryRow("SELECT Password FROM UserLogin WHERE Username = @username", sql.Named("username", username)).Scan(&storedPassword)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the stored password
	if password == storedPassword {
		fmt.Fprintln(w, "Login successful")
	} else {
		http.Error(w, "Login failed", http.StatusUnauthorized)
	}
}
