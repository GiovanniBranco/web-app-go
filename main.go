package main

import (
	"net/http"
	"text/template"

	"github.com/web-app-go/domain/models/products"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []products.Product{{
		Name:        "Shirt",
		Description: "Blue shirt",
		Price:       19.9,
		Amount:      10,
	},
		{
			Name:        "Shoe",
			Description: "Nice shoe",
			Price:       199.9,
			Amount:      15,
		},
	}

	templates.ExecuteTemplate(w, "Index", products)
}
