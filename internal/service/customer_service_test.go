package service

import (
	"testing"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
)

func TestCustomerService_Create_Success(t *testing.T) {

	// Arrange

	customerRepository := repository.NewMemoryCustomerRepository()

	customerService := NewCustomerService(customerRepository)

	// Act

	customer, err := customerService.Create(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"123456",
	)

	// Assert

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if customer == nil {
		t.Fatal("expected customer, got nil")
	}

	if customer.Name != "Ludimila" {
		t.Errorf(
			"expected name 'Ludimila', got '%s'",
			customer.Name,
		)
	}

	if customer.PasswordHash == "" {
		t.Fatal("expected password hash")
	}

	savedCustomer, err := customerRepository.FindByID(customer.ID)

	if err != nil {
		t.Fatalf("expected customer to be saved, got %v", err)
	}

	if savedCustomer.ID != customer.ID {
		t.Fatal("saved customer does not match created customer")
	}
}

func TestCustomerService_Create_DuplicateCPF(t *testing.T) {

	// Arrange

	customerRepository := repository.NewMemoryCustomerRepository()

	customerService := NewCustomerService(customerRepository)

	birthDate := time.Date(
		1995,
		5,
		20,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	_, err := customerService.Create(
		"Ludimila",
		"52998224725",
		birthDate,
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"123456",
	)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	customer, err := customerService.Create(
		"Maria",
		"52998224725",
		birthDate,
		"Rua B",
		"maria@email.com",
		"83888888888",
		"654321",
	)

	// Assert

	if err == nil {
		t.Fatal("expected duplicate CPF error")
	}

	if customer != nil {
		t.Fatal("expected nil customer")
	}
}
