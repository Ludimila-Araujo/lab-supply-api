package domain

import (
	"time"

	"github.com/google/uuid"
)

//representação do pedido feito pelo cliente

type Order struct {
	ID 			uuid.UUID
	Customer 	*Customer
	Items 		[]*OrderItem
	Status 		OrderStatus
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

func NewOrder(
	customer *Customer
) (*Order, error) {

	if customer == nil {
		return nil, ErrOrderCustomerRequired
	}

	now := time.Now()

	return &Order{

		ID: uuid.New(),
		Customer: customer,
		Items: []*OrderItem{},
		Status: OrderStatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

//função para adição de itens ao pedido

func (o *Order) AddItem(item *OrderItem) error {

	if item == nil {
		return ErrOrderItemRequired
	}

	if o.Status != OrderStatusPending {
		return ErrOrderCannotBeModified
	}

	o.Items = append(o.Items, item)
	o.Updated = time.Now()

	return nil
}

// função para cálculo do valor total do pedido

func (o *Order) Total() float64 {
	
	var total float64

	for _, item :=range o.Items{
		total += float64(item.Quantity) * item.UnitPrice
	}

	return total
}

