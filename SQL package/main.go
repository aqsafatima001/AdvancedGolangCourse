package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Data struct {
	id   int
	name string
}

func main() {

	connString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=learning;user id=Final_Year_Project;password=fyp;"

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

	result, err := db.Exec("INSERT INTO data (id, name) VALUES (5, 'Tomy');")
	CheckError(err)
	fmt.Println("Row inserted Sucessfully : ", result)
	// lastInsertedId, err := result.LastInsertId()
	// fmt.Println("Last Inserted Id = ", lastInsertedId)
	// CheckError(err)

	// rowsAffected, err := result.RowsAffected()
	// fmt.Println("Rows Affected = ", rowsAffected)
	// CheckError(err)

	rows, err := db.Query("Select * from data;")
	CheckError(err)
	fmt.Println("Rows Selected Successfully")

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		CheckError(err)
		fmt.Println(data)
	}

}

func CheckError(e error) {
	if e != nil {
		log.Fatalln("ERROR: ", e)
	}
}
