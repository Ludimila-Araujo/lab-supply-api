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
