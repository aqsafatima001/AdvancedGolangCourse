package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// product defined in the form of a structure
type product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

// Structure Declaration
var Products []product

// Main Func
func main() {

	Products = []product{
		product{Id: "1", Name: "Chair", Quantity: 100, Price: 140.7},
		product{Id: "2", Name: "Table", Quantity: 30, Price: 198.4},
	}
	HandleRequests()
}

// Function Definitions
func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/product", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", returnProduct)
	http.ListenAndServe(":8080", myRouter)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("End Point Hits: HOMEPAGE")
	fmt.Fprintf(w, "Welcome to home Page")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("End Point Hits: Return all products")
	json.NewEncoder(w).Encode(Products)
}

func returnProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, product := range Products {
		if product.Id == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}
