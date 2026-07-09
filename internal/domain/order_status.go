package domain

//representação do estado atual do pedido

type OrderStatus string

const (
	OrderStatusPending  OrderStatus = "PENDING"
	OrderStatusPaid     OrderStatus = "PAID"
	OrderStatusCanceled OrderStatus = "CANCELED"
)
