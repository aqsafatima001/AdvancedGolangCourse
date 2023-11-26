package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialise() error {
	connectionString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=Inventory;user id=Final_Year_Project;password=fyp;"

	var err error
	app.DB, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err.Error())
		return err
	}
	fmt.Println("Connected to the database!")

	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
	return nil

}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {

	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {

	error_message := map[string]string{"error": err}
	sendResponse(w, statusCode, error_message)
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {

	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)

}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	// app.Router.HandleFunc("/products/{id}", app.getProduct).Methods("GET")
}
