package repository

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"

	"github.com/google/uuid"
)

// responsável por armazenar dados do produto

type MemoryProductRepository struct {
	products map[uuid.UUID]*domain.Product
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]*domain.Product),
	}
}
