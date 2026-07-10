package repository

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

//repositório do produto define a persistência das operações

type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id uuid.UUID) (*domain.Product, error)
	FindAll() ([]*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uuid.UUID) error
}
