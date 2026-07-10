package repository

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

//Armazenamento de pedidos em memória

type MemoryOrderRepository struct {
	orders map[uuid.UUID]*domain.Order
}

//Criação in memory do pedido no repositório

func NewMemoryOrderRepository() *MemoryOrderRepository {

	return &MemoryOrderRepository{
		orders: make(map[uuid.UUID]*domain.Order),
	}
}

//Create

func (r *MemoryOrderRepository) Create(order *domain.Order) error {

	if _, exists := r.orders[order.ID]; exists {
		return ErrOrderAlreadyExists
	}

	r.orders[order.ID] = order

	return nil
}

//Retorno de um pedido de acordo com seu ID

func (r *MemoryOrderRepository) FindByID(id uuid.UUID) (*domain.Order, error) {

	order, exists := r.orders[id]

	if !exists {
		return nil, ErrOrderNotFound
	}

	return order, nil
}

// Atualização d eum pedido já existente
func (r *MemoryOrderRepository) Update(order *domain.Order) error {

	if _, exists := r.orders[order.ID]; !exists {
		return ErrOrderNotFound
	}

	r.orders[order.ID] = order

	return nil
}

// Remoção de um pedido da memória
func (r *MemoryOrderRepository) Delete(id uuid.UUID) error {

	if _, exists := r.orders[id]; !exists {
		return ErrOrderNotFound
	}

	delete(r.orders, id)

	return nil
}
