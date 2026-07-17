package repository

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	Create(customer *domain.Customer) error
	FindByID(id uuid.UUID) (*domain.Customer, error)
	FindByCPF(cpf string) (*domain.Customer, error)
	FindAll() ([]*domain.Customer, error)
	Update(customer *domain.Customer) error
	Delete(id uuid.UUID) error
}
