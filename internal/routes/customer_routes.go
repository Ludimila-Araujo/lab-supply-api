package routes

import (
	"net/http"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/controllers"
)

func RegisterCustomerRoutes(
	mux *http.ServeMux,
	controller *controllers.CustomerController,
) {

	mux.HandleFunc(
		"POST /customers",
		controller.Create,
	)

	mux.HandleFunc(
		"GET /customers",
		controller.FindAll,
	)

	mux.HandleFunc(
		"GET /customers/{id}",
		controller.FindByID,
	)
}
