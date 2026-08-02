package service

import (
	"errors"
	"testing"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
	"github.com/google/uuid"
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

func TestCustomerService_FindByID_Success(t *testing.T) {

	// Arrange

	customerRepository := repository.NewMemoryCustomerRepository()

	customerService := NewCustomerService(customerRepository)

	customer, err := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := customerRepository.Create(customer); err != nil {
		t.Fatal(err)
	}

	// Act

	foundCustomer, err := customerService.FindByID(customer.ID)

	// Assert

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if foundCustomer == nil {
		t.Fatal("expected customer, got nil")
	}

	if foundCustomer.ID != customer.ID {
		t.Error("expected same customer")
	}
}

func TestCustomerService_FindByID_NotFound(t *testing.T) {

	customerRepository := repository.NewMemoryCustomerRepository()

	customerService := NewCustomerService(customerRepository)

	id := uuid.New()

	customer, err := customerService.FindByID(id)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrCustomerNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrCustomerNotFound,
			err,
		)
	}

	if customer != nil {
		t.Fatal("expected nil customer")
	}
}

func TestCustomerService_FindAll(t *testing.T) {

	// Arrange

	customerRepository := repository.NewMemoryCustomerRepository()

	customerService := NewCustomerService(customerRepository)

	customer1, err := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	customer2, err := domain.NewCustomer(
		"Maria",
		"12345678901",
		time.Date(1990, 10, 10, 0, 0, 0, 0, time.UTC),
		"Rua B",
		"maria@email.com",
		"83888888888",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := customerRepository.Create(customer1); err != nil {
		t.Fatal(err)
	}

	if err := customerRepository.Create(customer2); err != nil {
		t.Fatal(err)
	}

	// Act

	customers, err := customerService.FindAll()

	// Assert

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(customers) != 2 {
		t.Fatalf("expected 2 customers, got %d", len(customers))
	}
}

func TestCustomerService_Update_Success(t *testing.T) {

	// Arrange

	customerRepository := repository.NewMemoryCustomerRepository()

	customerService := NewCustomerService(customerRepository)

	customer, err := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := customerRepository.Create(customer); err != nil {
		t.Fatal(err)
	}

	// Act

	customer.Name = "Ludimila Araújo"
	customer.Email = "novo@email.com"

	if err := customerService.Update(customer); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	// Assert

	updatedCustomer, err := customerRepository.FindByID(customer.ID)

	if err != nil {
		t.Fatal(err)
	}

	if updatedCustomer.Name != "Ludimila Araújo" {
		t.Errorf(
			"expected name 'Ludimila Araújo', got '%s'",
			updatedCustomer.Name,
		)
	}

	if updatedCustomer.Email != "novo@email.com" {
		t.Errorf(
			"expected email 'novo@email.com', got '%s'",
			updatedCustomer.Email,
		)
	}
}

func TestCustomerService_Update_CustomerNotFound(t *testing.T) {

	customerRepository := repository.NewMemoryCustomerRepository()

	customerService := NewCustomerService(customerRepository)

	customer, err := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	err = customerService.Update(customer)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrCustomerNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrCustomerNotFound,
			err,
		)
	}
}
