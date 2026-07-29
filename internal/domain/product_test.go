package domain

import (
	"errors"  //para comparação de erros
	"testing" //frameworking para testes

	"github.com/google/uuid"
)

func TestNewProduct(t *testing.T) {

	tests := []struct {
		name        string
		productName string
		description string
		brand       string
		price       float64
		stock       int
		expectedErr error
	}{
		{
			name:        "valid product",
			productName: "Micropipeta",
			description: "Micropipeta P20",
			brand:       "Eppendorf",
			price:       350.50,
			stock:       10,
			expectedErr: nil,
		},
		{
			name:        "empty product name",
			productName: "",
			description: "Micropipeta",
			brand:       "Eppendorf",
			price:       350,
			stock:       10,
			expectedErr: ErrProductNameRequired,
		},
		{
			name:        "blank prodcut name",
			productName: "      ",
			description: "Micropipeta",
			brand:       "Eppendorf",
			price:       350,
			stock:       10,
			expectedErr: ErrProductNameRequired,
		},
		{
			name:        "empty description",
			productName: "Micropipeta",
			description: "",
			brand:       "Eppendorf",
			price:       350.50,
			stock:       10,
			expectedErr: ErrProductDescriptionRequired,
		},
		{
			name:        "empty brand",
			productName: "Micropipeta",
			description: "Micropipeta P20",
			brand:       "",
			price:       350.50,
			stock:       10,
			expectedErr: ErrProductBrandRequired,
		},
		{
			name:        "price equals zero",
			productName: "Micropipeta",
			description: "Micropipeta P20",
			brand:       "Eppendorf",
			price:       0,
			stock:       10,
			expectedErr: ErrProductPriceRequired,
		},
		{
			name:        "negative price",
			productName: "Micropipeta",
			description: "Micropipeta P20",
			brand:       "Eppendorf",
			price:       -10,
			stock:       10,
			expectedErr: ErrProductPriceRequired,
		},
		{
			name:        "negative stock",
			productName: "Micropipeta",
			description: "Micropipeta P20",
			brand:       "Eppendorf",
			price:       350.50,
			stock:       -1,
			expectedErr: ErrProductStockRequired,
		},
		{
			name:        "zero stock is valid",
			productName: "Micropipeta",
			description: "Micropipeta P20",
			brand:       "Eppendorf",
			price:       350.50,
			stock:       0,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			//arrange:

			//act

			product, err := NewProduct(
				tt.productName,
				tt.description,
				tt.brand,
				tt.price,
				tt.stock,
			)

			//assert

			if !errors.Is(err, tt.expectedErr) {

				t.Fatalf(
					"expected error %v, got %v",
					tt.expectedErr,
					err,
				)
			}

			if tt.expectedErr != nil {
				return
			}

			if product == nil {
				t.Fatal("expected product, got nil")
			}

			if product.ID == uuid.Nil {
				t.Error("expected generated UUID")
			}

			if product.CreatedAt.IsZero() {
				t.Error("CreatedAt should be inicialized")
			}

			if product.UpdatedAt.IsZero() {
				t.Error("UpdatedAt should be inicialized")
			}

			if !product.CreatedAt.Equal(product.UpdatedAt) {
				t.Error("timestamps should match on creation")
			}

			if product.Name != tt.productName {
				t.Errorf("expected name %q, got %q", tt.productName, product.Name)
			}

			if product.Description != tt.description {
				t.Errorf("expected description %q, got %q", tt.description, product.Description)
			}

			if product.Brand != tt.brand {
				t.Errorf("expected brand %q, got %q", tt.brand, product.Brand)
			}

			if product.Price != tt.price {
				t.Errorf("expected price %f, got %f", tt.price, product.Price)
			}

			if product.Stock != tt.stock {
				t.Errorf("expected stock %d, got %d", tt.stock, product.Stock)
			}
		})

	}
}
