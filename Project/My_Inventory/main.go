package main

import (
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	app := App{}
	app.Initialise()
	app.Run(":8080")
}
