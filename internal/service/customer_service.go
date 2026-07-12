package service

import (
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// responsável pelas regras de negócio dos clientes.
type CustomerService struct {
	customerRepository repository.CustomerRepository
}

// Construtor.
func NewCustomerService(
	customerRepository repository.CustomerRepository,
) *CustomerService {

	return &CustomerService{
		customerRepository: customerRepository,
	}
}

// create

func (s *CustomerService) Create(
	name string,
	cpf string,
	birthDate time.Time,
	address string,
	email string,
	phone string,
	password string,
) (*domain.Customer, error) {

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	customer, err := domain.NewCustomer(
		name,
		cpf,
		birthDate,
		address,
		email,
		phone,
		string(passwordHash),
	)

	if err != nil {
		return nil, err
	}

	err = s.customerRepository.Create(customer)

	if err != nil {
		return nil, err
	}

	return customer, nil

}

// find id:

func (s *CustomerService) FindByID(
	id uuid.UUID,
) (*domain.Customer, error) {

	return s.customerRepository.FindByID(id)
}

// find all

func (s *CustomerService) FindAll() ([]*domain.Customer, error) {

	return s.customerRepository.FindAll()
}

// update

func (s *CustomerService) Update(
	customer *domain.Customer,
) error {

	return s.customerRepository.Update(customer)
}

//delete

func (s *CustomerService) Delete(
	id uuid.UUID,
) error {

	return s.customerRepository.Delete(id)
}
