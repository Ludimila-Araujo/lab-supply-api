package repository

import (
	"database/sql"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

//métodos para manipulação de dados

const (
	insertProductQuery = `
	INSERT INTO products (
	id,
	name,
	description,
	brand,
	price,
	stock,
	created_at,
	updated_at
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
`

	findProductByIDQuery = `
SELECT
	id,
	name,
	description,
	brand,
	price,
	stock,
	created_at,
	updated_at
	FROM products
	WHERE id = $1
`

	findAllProductsQuery = `

SELECT
		id,
		name,
		description,
		brand,
		price,
		stock,
		created_at,
		updated_at
	FROM products
	ORDER BY name

`

	updateProductQuery = `
	UPDATE products
	SET
		name = $1,
		description = $2,
		brand = $3,
		price = $4,
		stock = $5,
		updated_at = $6
	WHERE id = $7
	`

	deleteProductQuery = `
	DELETE FROM products
	WHERE id = $1
	`
)

const findProductByNameQuery = `
SELECT
	id,
	name,
	description,
	brand,
	price,
	stock,
	created_at,
	updated_at
FROM products
WHERE name = $1
`

// PostgresProductRepository implementa o ProductRepository utilizando PostgreSQL.

type PostgresProductRepository struct {
	db *sql.DB
}

// NewPostgresProductRepository cria uma nova instância do repositório PostgreSQL.

func NewPostgresProductRepository(
	db *sql.DB,
) *PostgresProductRepository {

	return &PostgresProductRepository{
		db: db,
	}
}

//método create para envio de parâmteros para o banco de dados:

func (r *PostgresProductRepository) Create(
	product *domain.Product,
) error {

	_, err := r.db.Exec(
		insertProductQuery,
		product.ID,
		product.Name,
		product.Description,
		product.Brand,
		product.Price,
		product.Stock,
		product.CreatedAt,
		product.UpdatedAt,
	)

	return err
}

//método para busca por ID

func (r *PostgresProductRepository) FindByID(id uuid.UUID) (*domain.Product, error) {

	product := &domain.Product{}

	err := r.db.QueryRow(
		findProductByIDQuery,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Brand,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrProductNotFound
	}

	if err != nil {
		return nil, err
	}

	return product, nil
}

//método para buscar tudo:

func (r *PostgresProductRepository) FindAll() ([]*domain.Product, error) {

	rows, err := r.db.Query(findAllProductsQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {

		product := &domain.Product{}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Brand,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

//método para atualização

func (r *PostgresProductRepository) Update(product *domain.Product) error {

	result, err := r.db.Exec(
		updateProductQuery,
		product.Name,
		product.Description,
		product.Brand,
		product.Price,
		product.Stock,
		product.UpdatedAt,
		product.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}

//método para deleção:

func (r *PostgresProductRepository) Delete(id uuid.UUID) error {

	result, err := r.db.Exec(
		deleteProductQuery,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}

//método para busca por nome

func (r *PostgresProductRepository) FindByName(
	name string,
) (*domain.Product, error) {

	product := &domain.Product{}

	err := r.db.QueryRow(
		findProductByNameQuery,
		name,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Brand,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrProductNotFound
	}

	if err != nil {
		return nil, err
	}

	return product, nil
}
