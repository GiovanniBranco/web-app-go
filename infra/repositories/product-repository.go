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

func GetProductById(connection *sql.DB, id int) products.Product {
	product := products.Product{}

	row, err := connection.Query("SELECT * FROM products WHERE id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	for row.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := row.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Amount = amount
		product.Name = name
		product.Description = description
		product.Price = price
	}

	return product

}

func InsertProduct(connection *sql.DB, product products.Product) {
	insert, err := connection.Prepare("INSERT INTO products  (name, description, price, amount) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(product.Name, product.Description, product.Price, product.Amount)
}

func EditProduct(connection *sql.DB, product products.Product) {
	update, err := connection.Prepare("UPDATE products SET name = $1, description = $2, price = $3, amount = $4 WHERE id = $5")

	if err != nil {
		panic(err.Error())
	}

	update.Exec(product.Name, product.Description, product.Price, product.Amount, product.Id)
}

func DeleteProduct(connection *sql.DB, id int) {
	delete, err := connection.Prepare("DELETE FROM products WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
}
