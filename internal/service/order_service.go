package service

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
)

//OrderService coordena os casos de uso relacionados aos pedidos.

type OrderService struct {
	productRepository repository.ProductRepository
	orderRepository   repository.OrderRepository
}

//construtor

func NewOrderService(
	productRepository repository.ProductRepository,
	orderRepository repository.OrderRepository,
) *OrderService {

	return &OrderService{
		productRepository: productRepository,
		orderRepository:   orderRepository,
	}
}

// CreateOrder cria um novo pedido para um cliente.

func (s *OrderService) CreateOrder(
	customer *domain.Customer,
	items []CreateOrderItemRequest,
) (*domain.Order, error) {

	order, err := domain.NewOrder(customer)
	if err != nil {
		return nil, err
	}

	for _, item := range items {

		if item.Quantity > item.Product.Stock {
			return nil, domain.ErrProductInsufficientStock
		}

		orderItem, err := domain.NewOrderItem(
			item.Product,
			item.Quantity,
		)

		if err != nil {
			return nil, err
		}

		if err := order.AddItem(orderItem); err != nil {
			return nil, err
		}

		item.Product.Stock -= item.Quantity

		if err := s.productRepository.Update(item.Product); err != nil {
			return nil, err
		}
	}

	if err := s.orderRepository.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}
