package service

import (
	"errors"
	"testing"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
	"github.com/google/uuid"
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

func TestProductService_FindByID_Success(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

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

	err = productRepository.Create(product)
	if err != nil {
		t.Fatal(err)
	}

	foundProduct, err := productService.FindByID(product.ID)

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

func TestProductService_FindByID_NotFound(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

	id := uuid.New()

	product, err := productService.FindByID(id)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, repository.ErrProductNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			repository.ErrProductNotFound,
			err,
		)
	}

	if product != nil {
		t.Fatal("expected nil product")
	}
}

func TestProductService_FindAll(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

	product1, err := domain.NewProduct(
		"Micropipeta",
		"Micropipeta P20",
		"Eppendorf",
		250,
		10,
	)

	if err != nil {
		t.Fatal(err)
	}

	product2, err := domain.NewProduct(
		"Centrífuga",
		"Centrífuga Digital",
		"Kasvi",
		1500,
		5,
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := productRepository.Create(product1); err != nil {
		t.Fatal(err)
	}

	if err := productRepository.Create(product2); err != nil {
		t.Fatal(err)
	}

	products, err := productService.FindAll()

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(products) != 2 {
		t.Fatalf("expected 2 products, got %d", len(products))
	}
}

func TestProductService_Update_Success(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

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

	product.Price = 500
	product.Stock = 20

	if err := productService.Update(product); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	updatedProduct, err := productRepository.FindByID(product.ID)
	if err != nil {
		t.Fatal(err)
	}

	if updatedProduct.Price != 500 {
		t.Errorf("expected price 500, got %.2f", updatedProduct.Price)
	}

	if updatedProduct.Stock != 20 {
		t.Errorf("expected stock 20, got %d", updatedProduct.Stock)
	}
}

func TestProductService_Update_ProductNotFound(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

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

	err = productService.Update(product)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, repository.ErrProductNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			repository.ErrProductNotFound,
			err,
		)
	}
}

func TestProductService_Delete_Success(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

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

	if err := productService.Delete(product.ID); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	_, err = productRepository.FindByID(product.ID)

	if !errors.Is(err, repository.ErrProductNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			repository.ErrProductNotFound,
			err,
		)
	}
}

func TestProductService_Delete_ProductNotFound(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()

	productService := NewProductService(productRepository)

	err := productService.Delete(uuid.New())

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, repository.ErrProductNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			repository.ErrProductNotFound,
			err,
		)
	}
}
