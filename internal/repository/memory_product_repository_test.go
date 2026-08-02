package repository

import (
	"errors"
	"testing"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

func TestMemoryProductRepository_Create_Success(t *testing.T) {

	// Arrange

	productRepository := NewMemoryProductRepository()

	product, err := domain.NewProduct(
		"Micropipeta",
		"Micropipeta P20",
		"Eppendorf",
		250,
		10,
	)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	err = productRepository.Create(product)

	// Assert

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	savedProduct, err := productRepository.FindByID(product.ID)

	if err != nil {
		t.Fatal(err)
	}

	if savedProduct.ID != product.ID {
		t.Error("expected same product")
	}
}

func TestMemoryProductRepository_Create_Duplicate(t *testing.T) {

	// Arrange

	productRepository := NewMemoryProductRepository()

	product, err := domain.NewProduct(
		"Micropipeta",
		"Micropipeta P20",
		"Eppendorf",
		250,
		10,
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := productRepository.Create(product); err != nil {
		t.Fatal(err)
	}

	// Act

	err = productRepository.Create(product)

	// Assert

	if err == nil {
		t.Fatal("expected error")
	}

	if err != domain.ErrProductAlreadyExists {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrProductAlreadyExists,
			err,
		)
	}
}

func TestMemoryProductRepository_FindByID_Success(t *testing.T) {

	productRepository := NewMemoryProductRepository()

	product, err := domain.NewProduct(
		"Micropipeta",
		"Descrição",
		"Eppendorf",
		250,
		10,
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := productRepository.Create(product); err != nil {
		t.Fatal(err)
	}

	foundProduct, err := productRepository.FindByID(product.ID)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if foundProduct == nil {
		t.Fatal("expected product, got nil")
	}

	if foundProduct.ID != product.ID {
		t.Error("expected same product")
	}
}

func TestMemoryProductRepository_FindByID_NotFound(t *testing.T) {

	productRepository := NewMemoryProductRepository()

	product, err := productRepository.FindByID(uuid.New())

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrProductNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrProductNotFound,
			err,
		)
	}

	if product != nil {
		t.Fatal("expected nil product")
	}
}

func TestMemoryProductRepository_FindAll(t *testing.T) {

	productRepository := NewMemoryProductRepository()

	product1, _ := domain.NewProduct(
		"Micropipeta",
		"Descrição",
		"Eppendorf",
		250,
		10,
	)

	product2, _ := domain.NewProduct(
		"Centrífuga",
		"Descrição",
		"Kasvi",
		1500,
		5,
	)

	productRepository.Create(product1)
	productRepository.Create(product2)

	products, err := productRepository.FindAll()

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(products) != 2 {
		t.Fatalf(
			"expected 2 products, got %d",
			len(products),
		)
	}
}
