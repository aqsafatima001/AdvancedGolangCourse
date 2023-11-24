package main

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialise() error {
	connectionString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=learning;user id=Final_Year_Project;password=fyp;"

	app.DB, err := sql.Open("sqlserver", connectionString)
}
