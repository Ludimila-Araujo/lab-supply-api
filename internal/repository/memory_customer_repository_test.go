package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

func TestMemoryCustomerRepository_Create_Success(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

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

	err = customerRepository.Create(customer)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	savedCustomer, err := customerRepository.FindByID(customer.ID)

	if err != nil {
		t.Fatal(err)
	}

	if savedCustomer.ID != customer.ID {
		t.Error("expected same customer")
	}
}

func TestMemoryCustomerRepository_FindByID_Success(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	customerRepository.Create(customer)

	foundCustomer, err := customerRepository.FindByID(customer.ID)

	if err != nil {
		t.Fatal(err)
	}

	if foundCustomer.ID != customer.ID {
		t.Error("expected same customer")
	}
}

func TestMemoryCustomerRepository_FindByID_NotFound(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	customer, err := customerRepository.FindByID(uuid.New())

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

func TestMemoryCustomerRepository_FindByCPF_Success(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	customerRepository.Create(customer)

	foundCustomer, err := customerRepository.FindByCPF("52998224725")

	if err != nil {
		t.Fatal(err)
	}

	if foundCustomer.CPF != "52998224725" {
		t.Error("expected same CPF")
	}
}

func TestMemoryCustomerRepository_FindByCPF_NotFound(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	customer, err := customerRepository.FindByCPF("11111111111")

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

func TestMemoryCustomerRepository_FindAll(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	customer1, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	customer2, _ := domain.NewCustomer(
		"Maria",
		"12345678901",
		time.Date(1990, 10, 10, 0, 0, 0, 0, time.UTC),
		"Rua B",
		"maria@email.com",
		"83888888888",
		"hash",
	)

	customerRepository.Create(customer1)
	customerRepository.Create(customer2)

	customers, err := customerRepository.FindAll()

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(customers) != 2 {
		t.Fatalf(
			"expected 2 customers, got %d",
			len(customers),
		)
	}
}

func TestMemoryCustomerRepository_Update_Success(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	customerRepository.Create(customer)

	customer.Name = "Ludimila Araújo"
	customer.Email = "novo@email.com"

	if err := customerRepository.Update(customer); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	updatedCustomer, err := customerRepository.FindByID(customer.ID)

	if err != nil {
		t.Fatal(err)
	}

	if updatedCustomer.Name != "Ludimila Araújo" {
		t.Errorf(
			"expected updated name, got %s",
			updatedCustomer.Name,
		)
	}

	if updatedCustomer.Email != "novo@email.com" {
		t.Errorf(
			"expected updated email, got %s",
			updatedCustomer.Email,
		)
	}
}

func TestMemoryCustomerRepository_Update_NotFound(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	err := customerRepository.Update(customer)

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

func TestMemoryCustomerRepository_Delete_Success(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	customerRepository.Create(customer)

	if err := customerRepository.Delete(customer.ID); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	_, err := customerRepository.FindByID(customer.ID)

	if !errors.Is(err, domain.ErrCustomerNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrCustomerNotFound,
			err,
		)
	}
}

func TestMemoryCustomerRepository_Delete_NotFound(t *testing.T) {

	customerRepository := NewMemoryCustomerRepository()

	err := customerRepository.Delete(uuid.New())

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
