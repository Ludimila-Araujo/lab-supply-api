package service

import "github.com/google/uuid"

type CreateOrderItemRequest struct {
	ProductID uuid.UUID
	Quantity  int
}
