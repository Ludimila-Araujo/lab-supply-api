package repository

import (
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"

	"github.com/google/uuid"
)

// responsável por armazenar dados do produto

type MemoryProductRepository struct {
	products map[uuid.UUID]*domain.Product
}

// função para criar uma nova instância do repositório de produtos em memória

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]*domain.Product),
	}
}

// função para criar um produto no repositório em memória

func (r *MemoryProductRepository) Create(product *domain.Product) error {

	if _, exists := r.products[product.ID]; exists {
		return ErrProductAlreadyExists
	}

	r.products[product.ID] = product

	return nil
}

// função para buscar um produto pelo ID no repositório em memória

func (r *MemoryProductRepository) FindByID(id uuid.UUID) (*domain.Product, error) {

	product, exists := r.products[id]

	if !exists {
		return nil, ErrProductNotFound
	}

	return product, nil
}

//função para retornar todos os produtos em estoque

func (r *MemoryProductRepository) FindAll() ([]*domain.Product, error) {

	products := make([]*domain.Product, 0, len(r.products))

	for _, product := range r.products {
		products = append(products, product)
	}

	return products, nil
}

// função para atualizar um produto no repositório em memória

func (r *MemoryProductRepository) Update(product *domain.Product) error {

	if _, exists := r.products[product.ID]; !exists {
		return ErrProductNotFound
	}

	r.products[product.ID] = product

	return nil
}

// função para deletar um produto no repositório em memória

func (r *MemoryProductRepository) Delete(id uuid.UUID) error {

	if _, exists := r.products[id]; !exists {
		return ErrProductNotFound
	}

	delete(r.products, id)

	return nil
}
