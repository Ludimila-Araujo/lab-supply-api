package domain

import (
	"testing"
	"time"
)

func mustParseDate(value string) time.Time {

	date, err := time.Parse(
		"2006-01-02",
		value,
	)

	if err != nil {
		panic(err)
	}

	return date
}

func TestNewOrder_Success(t *testing.T) {

	customer, err := NewCustomer(
		"Ludimila",
		"52998224725",
		mustParseDate("1995-05-20"),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	order, err := NewOrder(customer)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if order == nil {
		t.Fatal("expected order")
	}

	if order.Customer != customer {
		t.Fatal("customer mismatch")
	}

	if order.Status != OrderStatusPending {
		t.Fatalf(
			"expected %s, got %s",
			OrderStatusPending,
			order.Status,
		)
	}

	if len(order.Items) != 0 {
		t.Fatal("new order should start empty")
	}
}

func TestNewOrder_NilCustomer(t *testing.T) {

	order, err := NewOrder(nil)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrOrderCustomerRequired {
		t.Fatalf(
			"expected %v, got %v",
			ErrOrderCustomerRequired,
			err,
		)
	}

	if order != nil {
		t.Fatal("expected nil order")
	}
}

func TestOrder_AddItem_Success(t *testing.T) {

	// Arrange

	customer, err := NewCustomer(
		"Ludimila",
		"52998224725",
		mustParseDate("1995-05-20"),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	order, err := NewOrder(customer)

	if err != nil {
		t.Fatal(err)
	}

	product, err := NewProduct(
		"Micropipeta",
		"Micropipeta P20",
		"Eppendorf",
		250,
		10,
	)

	if err != nil {
		t.Fatal(err)
	}

	item, err := NewOrderItem(
		product,
		2,
	)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	err = order.AddItem(item)

	// Assert

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(order.Items) != 1 {
		t.Fatalf(
			"expected 1 item, got %d",
			len(order.Items),
		)
	}

	if order.Items[0] != item {
		t.Fatal("item was not added correctly")
	}
}

func TestOrder_AddItem_NilItem(t *testing.T) {

	// Arrange

	customer, err := NewCustomer(
		"Ludimila",
		"52998224725",
		mustParseDate("1995-05-20"),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	order, err := NewOrder(customer)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	err = order.AddItem(nil)

	// Assert

	if err == nil {
		t.Fatal("expected error")
	}

	if err != ErrOrderItemRequired {
		t.Fatalf(
			"expected %v, got %v",
			ErrOrderItemRequired,
			err,
		)
	}
}

func TestOrder_AddItem_OrderPaid(t *testing.T) {

	// Arrange

	customer, err := NewCustomer(
		"Ludimila",
		"52998224725",
		mustParseDate("1995-05-20"),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	order, err := NewOrder(customer)

	if err != nil {
		t.Fatal(err)
	}

	product, err := NewProduct(
		"Micropipeta",
		"Micropipeta P20",
		"Eppendorf",
		250,
		10,
	)

	if err != nil {
		t.Fatal(err)
	}

	item, err := NewOrderItem(
		product,
		2,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = order.Pay()

	if err != nil {
		t.Fatal(err)
	}

	// Act

	err = order.AddItem(item)

	// Assert

	if err == nil {
		t.Fatal("expected error")
	}

	if err != ErrOrderCannotBeModified {
		t.Fatalf(
			"expected %v, got %v",
			ErrOrderCannotBeModified,
			err,
		)
	}
}

func TestOrder_Total(t *testing.T) {

	// Arrange

	customer, err := NewCustomer(
		"Ludimila",
		"52998224725",
		mustParseDate("1995-05-20"),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	order, err := NewOrder(customer)

	if err != nil {
		t.Fatal(err)
	}

	product1, err := NewProduct(
		"Micropipeta",
		"Micropipeta P20",
		"Eppendorf",
		250,
		10,
	)

	if err != nil {
		t.Fatal(err)
	}

	product2, err := NewProduct(
		"Ponteira",
		"Ponteira Azul",
		"Kasvi",
		50,
		20,
	)

	if err != nil {
		t.Fatal(err)
	}

	item1, err := NewOrderItem(product1, 2)
	if err != nil {
		t.Fatal(err)
	}

	item2, err := NewOrderItem(product2, 3)
	if err != nil {
		t.Fatal(err)
	}

	err = order.AddItem(item1)
	if err != nil {
		t.Fatal(err)
	}

	err = order.AddItem(item2)
	if err != nil {
		t.Fatal(err)
	}

	// Act

	total := order.Total()

	// Assert

	expected := 650.0 // (2 x 250) + (3 x 50)

	if total != expected {
		t.Fatalf(
			"expected %.2f, got %.2f",
			expected,
			total,
		)
	}
}

func TestOrder_Pay_Success(t *testing.T) {

	// Arrange

	customer, err := NewCustomer(
		"Ludimila",
		"52998224725",
		mustParseDate("1995-05-20"),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	order, err := NewOrder(customer)

	if err != nil {
		t.Fatal(err)
	}

	// Act

	err = order.Pay()

	// Assert

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if order.Status != OrderStatusPaid {
		t.Fatalf(
			"expected %s, got %s",
			OrderStatusPaid,
			order.Status,
		)
	}
}

func TestOrder_Pay_InvalidStatus(t *testing.T) {

	customer, err := NewCustomer(
		"Ludimila",
		"52998224725",
		mustParseDate("1995-05-20"),
		"Rua A",
		"ludi@email.com",
		"83999999999",
		"hash",
	)

	if err != nil {
		t.Fatal(err)
	}

	order, err := NewOrder(customer)

	if err != nil {
		t.Fatal(err)
	}

	err = order.Pay()

	if err != nil {
		t.Fatal(err)
	}

	// Act

	err = order.Pay()

	// Assert

	if err == nil {
		t.Fatal("expected error")
	}

	if err != ErrOrderInvalidStatus {
		t.Fatalf(
			"expected %v, got %v",
			ErrOrderInvalidStatus,
			err,
		)
	}
}
