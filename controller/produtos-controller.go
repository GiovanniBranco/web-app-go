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

		if err != nil {
			log.Println("Error on convertion of price to float64")
		}

		amount, err := strconv.Atoi(r.FormValue("amount"))

		if err != nil {
			log.Println("Error on convertion of amount to int")
		}

		product := products.Product{Name: name, Price: price, Amount: amount, Description: description}
		repositories.InsertProduct(db, product)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	defer db.Close()
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := infra.GetConnection()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		log.Println("Error getting product id")
	}

	repositories.DeleteProduct(db, id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

	defer db.Close()
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	db := infra.GetConnection()
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		log.Println("Error getting product id")
	}

	product := repositories.GetProductById(db, id)

	templates.ExecuteTemplate(w, "Edit-Product", product)

	defer db.Close()
}

func UpdateProdut(w http.ResponseWriter, r *http.Request) {
	db := infra.GetConnection()
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))

		if err != nil {
			log.Println("Error getting product id")
		}

		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)

		if err != nil {
			log.Println("Error on convertion of price to float64")
		}

		amount, err := strconv.Atoi(r.FormValue("amount"))

		if err != nil {
			log.Println("Error on convertion of amount to int")
		}

		productUpdated := products.Product{Id: id, Name: name, Price: price, Amount: amount, Description: description}
		repositories.EditProduct(db, productUpdated)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	defer db.Close()
}
