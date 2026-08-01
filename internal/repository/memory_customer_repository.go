package repository

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]*domain.Customer
}

func NewMemoryCustomerRepository() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]*domain.Customer),
	}
}

func (r *MemoryCustomerRepository) Create(customer *domain.Customer) error {

	r.customers[customer.ID] = customer

	return nil
}

func (r *MemoryCustomerRepository) FindByID(
	id uuid.UUID,
) (*domain.Customer, error) {

	customer, exists := r.customers[id]

	if !exists {
		return nil, ErrCustomerNotFound
	}

	return customer, nil
}

func (r *MemoryCustomerRepository) FindByCPF(
	cpf string,
) (*domain.Customer, error) {

	for _, customer := range r.customers {

		if customer.CPF == cpf {
			return customer, nil
		}
	}

	return nil, ErrCustomerNotFound
}

func (r *MemoryCustomerRepository) FindAll() ([]*domain.Customer, error) {

	customers := make([]*domain.Customer, 0, len(r.customers))

	for _, customer := range r.customers {
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *MemoryCustomerRepository) Update(
	customer *domain.Customer,
) error {

	if _, exists := r.customers[customer.ID]; !exists {
		return ErrCustomerNotFound
	}

	r.customers[customer.ID] = customer

	return nil
}

func (r *MemoryCustomerRepository) Delete(
	id uuid.UUID,
) error {

	if _, exists := r.customers[id]; !exists {
		return ErrCustomerNotFound
	}

	delete(r.customers, id)

	return nil
}
