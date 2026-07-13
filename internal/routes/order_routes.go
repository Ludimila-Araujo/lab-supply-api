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
}
