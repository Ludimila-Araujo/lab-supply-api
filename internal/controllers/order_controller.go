package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/controllers/dto"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/service"

	"github.com/google/uuid"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(
	orderService *service.OrderService,
) *OrderController {

	return &OrderController{
		orderService: orderService,
	}
}

func (c *OrderController) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request dto.CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	customerID, err := uuid.Parse(request.CustomerID)
	if err != nil {
		http.Error(w, "invalid customer id", http.StatusBadRequest)
		return
	}

	items := make([]service.CreateOrderItemRequest, 0)

	for _, item := range request.Items {

		productID, err := uuid.Parse(item.ProductID)
		if err != nil {
			http.Error(w, "invalid product id", http.StatusBadRequest)
			return
		}

		items = append(items, service.CreateOrderItemRequest{
			ProductID: productID,
			Quantity:  item.Quantity,
		})
	}

	order, err := c.orderService.CreateOrder(
		customerID,
		items,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(order)
}
