package repository

import (
	"database/sql"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

// Repository responsável por persistir clientes
type PostgresCustomerRepository struct {
	db *sql.DB
}

// Construtor.
func NewPostgresCustomerRepository(
	db *sql.DB,
) *PostgresCustomerRepository {

	return &PostgresCustomerRepository{
		db: db,
	}
}

const (
	insertCustomerQuery = `
INSERT INTO customers (
	id,
	name,
	cpf,
	birth_date,
	address,
	email,
	phone,
	password_hash,
	created_at,
	updated_at
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
`

	findCustomerByIDQuery = `
SELECT
	id,
	name,
	cpf,
	birth_date,
	address,
	email,
	phone,
	password_hash,
	created_at,
	updated_at
FROM customers
WHERE id = $1
`

	findAllCustomersQuery = `
SELECT
	id,
	name,
	cpf,
	birth_date,
	address,
	email,
	phone,
	password_hash,
	created_at,
	updated_at
FROM customers
ORDER BY name
`

	updateCustomerQuery = `
UPDATE customers
SET
	name = $1,
	cpf = $2,
	birth_date = $3,
	address = $4,
	email = $5,
	phone = $6,
	password_hash = $7,
	updated_at = $8
WHERE id = $9
`

	deleteCustomerQuery = `
DELETE FROM customers
WHERE id = $1
`
)

// create:
func (r *PostgresCustomerRepository) Create(
	customer *domain.Customer,
) error {

	_, err := r.db.Exec(
		insertCustomerQuery,
		customer.ID,
		customer.Name,
		customer.CPF,
		customer.BirthDate,
		customer.Address,
		customer.Email,
		customer.Phone,
		customer.PasswordHash,
		customer.CreatedAt,
		customer.UpdatedAt,
	)

	return err
}

// find by ID:

func (r *PostgresCustomerRepository) FindByID(
	id uuid.UUID,
) (*domain.Customer, error) {

	customer := &domain.Customer{}

	err := r.db.QueryRow(
		findCustomerByIDQuery,
		id,
	).Scan(
		&customer.ID,
		&customer.Name,
		&customer.CPF,
		&customer.BirthDate,
		&customer.Address,
		&customer.Email,
		&customer.Phone,
		&customer.PasswordHash,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrCustomerNotFound
	}

	if err != nil {
		return nil, err
	}

	return customer, nil
}

// find all:

func (r *PostgresCustomerRepository) FindAll() ([]*domain.Customer, error) {

	rows, err := r.db.Query(findAllCustomersQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []*domain.Customer

	for rows.Next() {

		customer := &domain.Customer{}

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.CPF,
			&customer.BirthDate,
			&customer.Address,
			&customer.Email,
			&customer.Phone,
			&customer.PasswordHash,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

// update:

func (r *PostgresCustomerRepository) Update(
	customer *domain.Customer,
) error {

	result, err := r.db.Exec(
		updateCustomerQuery,
		customer.Name,
		customer.CPF,
		customer.BirthDate,
		customer.Address,
		customer.Email,
		customer.Phone,
		customer.PasswordHash,
		customer.UpdatedAt,
		customer.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrCustomerNotFound
	}

	return nil
}

// delete:

func (r *PostgresCustomerRepository) Delete(
	id uuid.UUID,
) error {

	result, err := r.db.Exec(
		deleteCustomerQuery,
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
		return ErrCustomerNotFound
	}

	return nil
}
