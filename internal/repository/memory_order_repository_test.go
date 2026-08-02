package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

func TestMemoryOrderRepository_Create_Success(t *testing.T) {

	orderRepository := NewMemoryOrderRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	order, _ := domain.NewOrder(customer)

	err := orderRepository.Create(order)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	savedOrder, err := orderRepository.FindByID(order.ID)

	if err != nil {
		t.Fatal(err)
	}

	if savedOrder.ID != order.ID {
		t.Error("expected same order")
	}
}

func TestMemoryOrderRepository_Create_Duplicate(t *testing.T) {

	orderRepository := NewMemoryOrderRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	order, _ := domain.NewOrder(customer)

	orderRepository.Create(order)

	err := orderRepository.Create(order)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrOrderAlreadyExists) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrOrderAlreadyExists,
			err,
		)
	}
}

func TestMemoryOrderRepository_FindByID_NotFound(t *testing.T) {

	orderRepository := NewMemoryOrderRepository()

	order, err := orderRepository.FindByID(uuid.New())

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrOrderNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrOrderNotFound,
			err,
		)
	}

	if order != nil {
		t.Fatal("expected nil order")
	}
}

func TestMemoryOrderRepository_FindAll(t *testing.T) {

	orderRepository := NewMemoryOrderRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	for i := 0; i < 5; i++ {

		order, _ := domain.NewOrder(customer)

		orderRepository.Create(order)
	}

	orders, err := orderRepository.FindAll(2, 1)

	if err != nil {
		t.Fatal(err)
	}

	if len(orders) != 2 {
		t.Fatalf(
			"expected 2 orders, got %d",
			len(orders),
		)
	}
}

func TestMemoryOrderRepository_RestoreStock(t *testing.T) {

	product, _ := domain.NewProduct(
		"Micropipeta",
		"Descrição",
		"Eppendorf",
		250,
		5,
	)

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	item, _ := domain.NewOrderItem(product, 3)

	order, _ := domain.NewOrder(customer)

	order.AddItem(item)

	orderRepository := NewMemoryOrderRepository()

	initialStock := product.Stock

	err := orderRepository.RestoreStock(order)

	if err != nil {
		t.Fatal(err)
	}

	if product.Stock != initialStock+3 {
		t.Fatalf(
			"expected stock %d, got %d",
			initialStock+3,
			product.Stock,
		)
	}
}

func TestMemoryOrderRepository_Update_Success(t *testing.T) {

	orderRepository := NewMemoryOrderRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	order, _ := domain.NewOrder(customer)

	if err := orderRepository.Create(order); err != nil {
		t.Fatal(err)
	}

	order.Status = domain.OrderStatusPaid

	if err := orderRepository.Update(order); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	updatedOrder, err := orderRepository.FindByID(order.ID)

	if err != nil {
		t.Fatal(err)
	}

	if updatedOrder.Status != domain.OrderStatusPaid {
		t.Fatalf(
			"expected status %s, got %s",
			domain.OrderStatusPaid,
			updatedOrder.Status,
		)
	}
}

func TestMemoryOrderRepository_Update_NotFound(t *testing.T) {

	orderRepository := NewMemoryOrderRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	order, _ := domain.NewOrder(customer)

	err := orderRepository.Update(order)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrOrderNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrOrderNotFound,
			err,
		)
	}
}

func TestMemoryOrderRepository_Delete_Success(t *testing.T) {

	orderRepository := NewMemoryOrderRepository()

	customer, _ := domain.NewCustomer(
		"Ludimila",
		"52998224725",
		time.Date(1995, 5, 20, 0, 0, 0, 0, time.UTC),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	order, _ := domain.NewOrder(customer)

	orderRepository.Create(order)

	if err := orderRepository.Delete(order.ID); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	_, err := orderRepository.FindByID(order.ID)

	if !errors.Is(err, domain.ErrOrderNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrOrderNotFound,
			err,
		)
	}
}

func TestMemoryOrderRepository_Delete_NotFound(t *testing.T) {

	orderRepository := NewMemoryOrderRepository()

	err := orderRepository.Delete(uuid.New())

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrOrderNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrOrderNotFound,
			err,
		)
	}
}
