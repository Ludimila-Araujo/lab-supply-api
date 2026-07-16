package routes

import (
	"net/http"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/controllers"
)

func RegisterOrderRoutes(
	mux *http.ServeMux,
	controller *controllers.OrderController,
) {

	mux.HandleFunc(
		"POST /orders",
		controller.Create,
	)

	mux.HandleFunc(
		"GET /orders",
		controller.FindAll,
	)

	mux.HandleFunc(
		"GET /orders/{id}",
		controller.FindByID,
	)

	mux.HandleFunc(
		"POST /orders/{id}/pay",
		controller.Pay,
	)
}
