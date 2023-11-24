package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", homepage)
	http.ListenAndServe(":8080", nil)

}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home Page")
}
