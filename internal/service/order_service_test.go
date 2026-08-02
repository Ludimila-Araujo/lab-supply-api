package service

import (
	"errors"
	"testing"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
	"github.com/google/uuid"
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

func TestOrderService_CreateOrder_CustomerNotFound(t *testing.T) {

	productRepository := repository.NewMemoryProductRepository()
	customerRepository := repository.NewMemoryCustomerRepository()
	orderRepository := repository.NewMemoryOrderRepository()

	orderService := NewOrderService(
		productRepository,
		customerRepository,
		orderRepository,
	)

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

	order, err := orderService.CreateOrder(
		uuid.New(),
		items,
	)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrCustomerNotFound) {
		t.Fatalf(
			"expected %v, got %v",
			repository.ErrCustomerNotFound,
			err,
		)
	}

	if order != nil {
		t.Fatal("expected nil order")
	}
}

func TestOrderService_CreateOrder_ProductNotFound(t *testing.T) {

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

	items := []CreateOrderItemRequest{
		{
			ProductID: uuid.New(), // Produto inexistente
			Quantity:  2,
		},
	}

	order, err := orderService.CreateOrder(
		customer.ID,
		items,
	)

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

	if order != nil {
		t.Fatal("expected nil order")
	}
}

func TestOrderService_CreateOrder_InsufficientStock(t *testing.T) {

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
		5, // estoque disponível
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
			Quantity:  10, // maior que o estoque
		},
	}

	order, err := orderService.CreateOrder(
		customer.ID,
		items,
	)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, domain.ErrProductInsufficientStock) {
		t.Fatalf(
			"expected %v, got %v",
			domain.ErrProductInsufficientStock,
			err,
		)
	}

	if order != nil {
		t.Fatal("expected nil order")
	}
}

func TestOrderService_Pay_Success(t *testing.T) {

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

	order, err := domain.NewOrder(customer)
	if err != nil {
		t.Fatal(err)
	}

	if err := orderRepository.Create(order); err != nil {
		t.Fatal(err)
	}

	if err := orderService.Pay(order.ID); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	savedOrder, err := orderRepository.FindByID(order.ID)
	if err != nil {
		t.Fatal(err)
	}

	if savedOrder.Status != domain.OrderStatusPaid {
		t.Fatalf(
			"expected status %s, got %s",
			domain.OrderStatusPaid,
			savedOrder.Status,
		)
	}
}

func TestOrderService_Cancel_Success(t *testing.T) {

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

	item, err := domain.NewOrderItem(product, 2)
	if err != nil {
		t.Fatal(err)
	}

	order, err := domain.NewOrder(customer)
	if err != nil {
		t.Fatal(err)
	}

	if err := order.AddItem(item); err != nil {
		t.Fatal(err)
	}

	if err := orderRepository.Create(order); err != nil {
		t.Fatal(err)
	}

	initialStock := product.Stock

	if err := orderService.Cancel(order.ID); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	savedOrder, err := orderRepository.FindByID(order.ID)
	if err != nil {
		t.Fatal(err)
	}

	if savedOrder.Status != domain.OrderStatusCanceled {
		t.Fatalf(
			"expected status %s, got %s",
			domain.OrderStatusCanceled,
			savedOrder.Status,
		)
	}

	if product.Stock != initialStock+2 {
		t.Fatalf(
			"expected stock %d, got %d",
			initialStock+2,
			product.Stock,
		)
	}
}
