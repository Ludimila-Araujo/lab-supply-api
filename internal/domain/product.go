package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

//representação de produto vendido pela loja

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Brand       string
	Price       float64
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//função para criação e validação de um novo produto

func NewProduct(
	name string,
	description string,
	brand string,
	price float64,
	stock int,
) (*Product, error) {

	if strings.TrimSpace(name) == "" {
		return nil, ErrProductNameRequired
	}

	if strings.TrimSpace(description) == "" {
		return nil, ErrProductDescriptionRequired
	}

	if strings.TrimSpace(brand) == "" {
		return nil, ErrProductBrandRequired
	}

	if price <= 0 {
		return nil, ErrProductPriceRequired
	}

	if stock < 0 {
		return nil, ErrProductStockRequired
	}

	now := time.Now()

	return &Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Brand:       brand,
		Price:       price,
		Stock:       stock,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
