package repositories

import (
	"database/sql"

	"github.com/web-app-go/domain/models/products"
)

func GetProducts(connection *sql.DB) []products.Product {
	productsList := []products.Product{}

	rows, err := connection.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {

		product := products.Product{}

		var id, amount int
		var name, description string
		var price float64

		err := rows.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Amount = amount
		product.Name = name
		product.Description = description
		product.Price = price

		productsList = append(productsList, product)
	}

	return productsList
}

func InsertProduct(connection *sql.DB, product products.Product) {
	insert, err := connection.Prepare("INSERT INTO products  (name, description, price, amount) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(product.Name, product.Description, product.Price, product.Amount)
}
