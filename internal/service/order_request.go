package service

import "github.com/Ludimila-Araujo/lab-supply-api/internal/domain"

//criação de uma reqiusição para criar um pedido

type CreateOrderItemRequest struct {
	Product  *domain.Product
	Quantity int
}
