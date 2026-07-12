package service

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"

	"github.com/google/uuid"
)

//struct

type ProductService struct {
	productRepository repository.ProductRepository
}

//construtor

func NewProductService(
	productRepository repository.ProductRepository,
) *ProductService {

	return &ProductService{
		productRepository: productRepository,
	}
}

//manipulação de dados - CRUD completo

func (s *ProductService) Create(product *domain.Product) error {
	return s.productRepository.Create(product)
}

func (s *ProductService) FindByID(id uuid.UUID) (*domain.Product, error) {
	return s.productRepository.FindByID(id)
}

func (s *ProductService) FindAll() ([]*domain.Product, error) {
	return s.productRepository.FindAll()
}

func (s *ProductService) Update(product *domain.Product) error {
	return s.productRepository.Update(product)
}

func (s *ProductService) Delete(id uuid.UUID) error {
	return s.productRepository.Delete(id)
}
