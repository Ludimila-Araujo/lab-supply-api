package repository

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	FindByID(id uuid.UUID) (*domain.Order, error)
	Update(order *domain.Order) error
	Delete(id uuid.UUID) error
}
