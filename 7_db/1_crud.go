// CREATE TABLE products (
//     product_id SERIAL PRIMARY KEY,
//     name TEXT NOT NULL,
//     price NUMERIC(10,2) NOT NULL,
//     category TEXT NOT NULL
// );

// go mod init simple-db

// go get github.com/lib/pq

package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

func main() {
	// Update connection info to match your DB
	connStr := "user=postgres password=password dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// CREATE
	// var id int
	// err = db.QueryRow(
	// 	"INSERT INTO products (name, price, category) VALUES ($1, $2, $3) RETURNING product_id",
	// 	"Pen", 9.99, "Stationery").Scan(&id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Inserted Product with ID: %d\n", id)

	// READ
	// var p Product
	// err = db.QueryRow(
	// 	"SELECT product_id, name, price, category FROM products WHERE product_id = $1", 3).
	// 	Scan(&p.ID, &p.Name, &p.Price, &p.Category)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Read Product: %+v\n", p)

	// UPDATE
	// _, err = db.Exec("UPDATE products SET price = $1 WHERE product_id = $2", 2.99, 3)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Updated Product price.")

	// READ ALL
	// rows, err := db.Query("SELECT product_id, name, price, category FROM products")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// fmt.Println("All products:")
	// for rows.Next() {
	// 	var prod Product
	// 	if err := rows.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Category); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%+v\n", prod)
	// }

	// DELETE
	// _, err = db.Exec("DELETE FROM products WHERE product_id = $1", 3)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Deleted Product.")
}
