package routes

import (
	"net/http"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/controllers"
)

func RegisterProductRoutes(
	mux *http.ServeMux,
	controller *controllers.ProductController,
) {

	mux.HandleFunc("POST /products", controller.Create)

	mux.HandleFunc("GET /products", controller.FindAll)

	mux.HandleFunc("GET /products/{id}", controller.FindByID)
}
