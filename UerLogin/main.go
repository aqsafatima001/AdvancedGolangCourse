package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func main() {
	connString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=learning;user id=Final_Year_Project;password=fyp;"
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err.Error())
		return
	}

	fmt.Println("Connected to the database!")

	http.HandleFunc("/", serveLogin)
	http.HandleFunc("/api/login", loginAPI)

	// Serve static files (CSS, JavaScript, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

// func serveLogin(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Serve the login page
// 	http.ServeFile(w, r, "templates/login.html")
// }

// func loginAPI(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Parse the form data
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
// 		return
// 	}

// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	// Query the database for the user's credentials
// 	var storedPassword string
// 	// err = db.QueryRow("SELECT Password FROM UserLogin WHERE Username = ?", username).Scan(&storedPassword)
// 	err = db.QueryRow("SELECT Password FROM UserLogin WHERE Username = @username", sql.Named("username", username)).Scan(&storedPassword)
// 	if err != nil {
// 		log.Printf("Error querying database: %v", err)
// 		http.Error(w, "User not found", http.StatusUnauthorized)
// 		return
// 	}

// 	// Compare the provided password with the stored password
// 	if password == storedPassword {
// 		fmt.Fprintln(w, "Login successful")
// 	} else {
// 		http.Error(w, "Login failed", http.StatusUnauthorized)
// 	}
// }
