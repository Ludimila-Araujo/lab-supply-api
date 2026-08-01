package service

import (
	"errors"
	"testing"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
)

func TestProductService_Create_Success(t *testing.T) {

	// Arrange

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

	// Act

	product, err := productService.Create(
		"Micropipeta",
		"Micropipeta P20",
		"Eppendorf",
		250.00,
		10,
	)

	// Assert

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if product == nil {
		t.Fatal("expected product, got nil")
	}

	if product.Name != "Micropipeta" {
		t.Errorf("expected product name 'Micropipeta', got '%s'", product.Name)
	}

	savedProduct, err := productRepository.FindByID(product.ID)

	if err != nil {
		t.Fatalf("expected product to be saved, got error %v", err)
	}

	if savedProduct.ID != product.ID {
		t.Error("saved product does not match created product")
	}

}

func TestProductService_Create_InvalidName(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

	product, err := productService.Create(
		"",
		"Micropipeta P20",
		"Eppendorf",
		250.00,
		10,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, domain.ErrProductNameRequired) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrProductNameRequired,
			err,
		)
	}

	if product != nil {
		t.Fatal("expected nil product")
	}

	products, err := productRepository.FindAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != 0 {
		t.Fatalf(
			"expected repository to be empty, got %d products",
			len(products),
		)
	}
}
