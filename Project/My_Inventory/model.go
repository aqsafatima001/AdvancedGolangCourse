package main

import (
	"database/sql"
	"fmt"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]product, error) {
	query := "Select id,name,quantity,price from products"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	products := []product{}
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)

	}
	return products, nil
}

func (p *product) getProduct(db *sql.DB) error {
	query := fmt.Sprintf("Select name,quantity,price FROM products where id = %v", p.ID)
	row := db.QueryRow(query)
	err := row.Scan(&p.Name, &p.Quantity, &p.Price)
	if err != nil {
		return err
	}
	return nil
}

func (p *product) createProduct(db *sql.DB) error {
	// Define the SQL query with named parameters and OUTPUT clause
	query := `
		INSERT INTO products (Name, Quantity, Price)
		OUTPUT INSERTED.ID
		VALUES (@Name, @Quantity, @Price)
	`

	// Prepare the SQL statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with named parameters
	var id int
	err = stmt.QueryRow(sql.Named("Name", p.Name), sql.Named("Quantity", p.Quantity), sql.Named("Price", p.Price)).Scan(&id)
	if err != nil {
		return err
	}

	p.ID = id
	return nil
}

func (p *product) updateProduct(db *sql.DB) error {
	query := fmt.Sprintf("Update products set name ='%v' , quanitiy = %v , price = %v where id = %v ", p.Name, p.Quantity, p.Price, p.ID)

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

}
