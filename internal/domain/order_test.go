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
