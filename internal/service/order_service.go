package service

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
)

//OOrderService coordena os casos de uso relacionados aos pedidos.

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

func (s *OrderService) CreateOrder(
	customer *domain.Customer,
	items []*domain.OrderItem,
) (*domain.Order, error) {

	order, err := domain.NewOrder(customer)
	if err != nil {
		return nil, err
	}

	for _, item := range items {

		if item.Quantity > item.Product.Stock {
			return nil, domain.ErrProductInsufficientStock
		}

		if err := order.AddItem(item); err != nil {
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
