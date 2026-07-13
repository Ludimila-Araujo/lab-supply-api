package service

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
	"github.com/google/uuid"
)

//OrderService coordena os casos de uso relacionados aos pedidos.

type OrderService struct {
	productRepository  repository.ProductRepository
	customerRepository repository.CustomerRepository
	orderRepository    repository.OrderRepository
}

//construtor

func NewOrderService(
	productRepository repository.ProductRepository,
	customerRepository repository.CustomerRepository,
	orderRepository repository.OrderRepository,
) *OrderService {

	return &OrderService{
		productRepository:  productRepository,
		customerRepository: customerRepository,
		orderRepository:    orderRepository,
	}
}

// CreateOrder cria um novo pedido para um cliente.

func (s *OrderService) CreateOrder(
	customerID uuid.UUID,
	items []CreateOrderItemRequest,
) (*domain.Order, error) {

	customer, err := s.customerRepository.FindByID(customerID)
	if err != nil {
		return nil, err
	}

	order, err := domain.NewOrder(customer)
	if err != nil {
		return nil, err
	}

	for _, item := range items {

		product, err := s.productRepository.FindByID(item.ProductID)
		if err != nil {
			return nil, err
		}

		if item.Quantity > product.Stock {
			return nil, domain.ErrProductInsufficientStock
		}

		orderItem, err := domain.NewOrderItem(
			product,
			item.Quantity,
		)

		if err != nil {
			return nil, err
		}

		if err := order.AddItem(orderItem); err != nil {
			return nil, err
		}
	}

	if err := s.orderRepository.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) FindAll(
	limit, offset int,
) ([]*domain.Order, error) {

	return s.orderRepository.FindAll(limit, offset)
}
