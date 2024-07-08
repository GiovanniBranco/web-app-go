package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/web-app-go/domain/models/products"
	"github.com/web-app-go/infra"
	"github.com/web-app-go/infra/repositories"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := infra.GetConnection()
	productsList := repositories.GetProducts(db)

	templates.ExecuteTemplate(w, "Index", productsList)
	defer db.Close()
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New-Product", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db := infra.GetConnection()
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		amount, err := strconv.Atoi(r.FormValue("amount"))

		if err != nil {
			log.Println("Error on convertion of price to float64 or amount to int")
		}

		product := products.Product{Name: name, Price: price, Amount: amount, Description: description}
		repositories.InsertProduct(db, product)
	}
	http.Redirect(w, r, "/", 301)
	defer db.Close()
}
