package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	connString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=PVFC;user id=pvfc;password=pvfc;"

	// Open a connection to the database
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err.Error())
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

}
