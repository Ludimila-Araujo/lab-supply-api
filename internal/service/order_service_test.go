package service

import (
	"testing"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
)

func TestOrderService_CreateOrder_Success(t *testing.T) {

	// Arrange

	productRepository := repository.NewMemoryProductRepository()
	customerRepository := repository.NewMemoryCustomerRepository()
	orderRepository := repository.NewMemoryOrderRepository()

	orderService := NewOrderService(
		productRepository,
		customerRepository,
		orderRepository,
	)

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

	items := []CreateOrderItemRequest{
		{
			ProductID: product.ID,
			Quantity:  2,
		},
	}

	// Act

	order, err := orderService.CreateOrder(
		customer.ID,
		items,
	)

	// Assert

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if order == nil {
		t.Fatal("expected order, got nil")
	}

	if order.Customer.ID != customer.ID {
		t.Error("customer mismatch")
	}

	if order.Status != domain.OrderStatusPending {
		t.Errorf(
			"expected status %s, got %s",
			domain.OrderStatusPending,
			order.Status,
		)
	}

	if len(order.Items) != 1 {
		t.Fatalf(
			"expected 1 item, got %d",
			len(order.Items),
		)
	}

	if order.Items[0].Quantity != 2 {
		t.Errorf(
			"expected quantity 2, got %d",
			order.Items[0].Quantity,
		)
	}

	savedOrder, err := orderRepository.FindByID(order.ID)

	if err != nil {
		t.Fatalf("expected saved order, got %v", err)
	}

	if savedOrder.ID != order.ID {
		t.Error("saved order mismatch")
	}
}
