package routes

import (
	"net/http"

	"github.com/web-app-go/controller"
)

func HandleRequests() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.NewProduct)
	http.HandleFunc("/insert", controller.InsertProduct)
	http.HandleFunc("/delete", controller.DeleteProduct)
	http.HandleFunc("/edit", controller.EditProduct)
	http.HandleFunc("/update", controller.UpdateProdut)
}
