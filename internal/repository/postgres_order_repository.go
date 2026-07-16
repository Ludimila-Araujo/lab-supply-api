package repository

import (
	"database/sql"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/google/uuid"
)

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(
	db *sql.DB,
) *PostgresOrderRepository {

	return &PostgresOrderRepository{
		db: db,
	}
}

const findOrderByIDQuery = `
SELECT 
	o.id,
	o.status,
	o.created_at,
	o.updated_at,

	c.id,
	c.name,
	c.cpf,
	c.birth_date,
	c.address,
	c.email,
	c.phone,
	c.password_hash,
	c.created_at,
	c.updated_at

FROM orders o
INNER JOIN customers c
	ON c.id = o.customer_id
WHERE o.id = $1
`

const findOrderItemsQuery = `
SELECT
	p.id,
	p.name,
	p.description,
	p.brand,
	p.price,
	p.stock,
	p.created_at,
	p.updated_at,

	oi.quantity,
	oi.unit_price

FROM order_items oi

INNER JOIN products p
	ON p.id = oi.product_id

WHERE oi.order_id = $1
`

const findAllOrdersQuery = `
	SELECT
		o.id,
		o.customer_id,
		o.status,
		o.created_at,
		o.updated_at
	
	FROM orders o

	ORDER BY o.created_at DESC

	LIMIT $1
	OFFSET $2
`

func (r *PostgresOrderRepository) Create(
	order *domain.Order,
) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	const insertOrderQuery = `
	INSERT INTO orders (
		id,
		customer_id,
		status,
		created_at,
		updated_at
	)
	VALUES ($1,$2,$3,$4,$5)
	`

	_, err = tx.Exec(
		insertOrderQuery,
		order.ID,
		order.Customer.ID,
		order.Status,
		order.CreatedAt,
		order.UpdatedAt,
	)

	if err != nil {
		return err
	}

	const insertOrderItemQuery = `
	INSERT INTO order_items (
		order_id,
		product_id,
		quantity,
		unit_price
	)
	VALUES ($1,$2,$3,$4)
	`

	const updateStockQuery = `
	UPDATE products
	SET
		stock = stock - $1,
		updated_at = NOW()
	WHERE id = $2
	`

	for _, item := range order.Items {

		_, err = tx.Exec(
			insertOrderItemQuery,
			order.ID,
			item.Product.ID,
			item.Quantity,
			item.UnitPrice,
		)

		if err != nil {
			return err
		}

		result, err := tx.Exec(
			updateStockQuery,
			item.Quantity,
			item.Product.ID,
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
	}

	return tx.Commit()
}

func (r *PostgresOrderRepository) loadOrderItems(
	order *domain.Order,
) error {

	rows, err := r.db.Query(
		findOrderItemsQuery,
		order.ID,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		product := &domain.Product{}
		item := &domain.OrderItem{
			Product: product,
		}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Brand,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,

			&item.Quantity,
			&item.UnitPrice,
		)

		if err != nil {
			return err
		}

		order.Items = append(order.Items, item)
	}

	return rows.Err()
}

func (r *PostgresOrderRepository) FindByID(
	id uuid.UUID,
) (*domain.Order, error) {

	order := &domain.Order{
		Customer: &domain.Customer{},
		Items:    []*domain.OrderItem{},
	}

	err := r.db.QueryRow(
		findOrderByIDQuery,
		id,
	).Scan(
		&order.ID,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,

		&order.Customer.ID,
		&order.Customer.Name,
		&order.Customer.CPF,
		&order.Customer.BirthDate,
		&order.Customer.Address,
		&order.Customer.Email,
		&order.Customer.Phone,
		&order.Customer.PasswordHash,
		&order.Customer.CreatedAt,
		&order.Customer.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrOrderNotFound
	}

	if err := r.loadOrderItems(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (r *PostgresOrderRepository) FindAll(
	limit, offset int,
) ([]*domain.Order, error) {

	rows, err := r.db.Query(
		findAllOrdersQuery,
		limit,
		offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := make([]*domain.Order, 0)

	for rows.Next() {

		order := &domain.Order{
			Customer: &domain.Customer{},
			Items:    []*domain.OrderItem{},
		}

		err := rows.Scan(
			&order.ID,
			&order.Customer.ID,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Busca os dados completos do cliente
		customer, err := r.db.Query(
			`
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
			`,
			order.Customer.ID,
		)

		if err != nil {
			return nil, err
		}

		if customer.Next() {

			err = customer.Scan(
				&order.Customer.ID,
				&order.Customer.Name,
				&order.Customer.CPF,
				&order.Customer.BirthDate,
				&order.Customer.Address,
				&order.Customer.Email,
				&order.Customer.Phone,
				&order.Customer.PasswordHash,
				&order.Customer.CreatedAt,
				&order.Customer.UpdatedAt,
			)

			if err != nil {
				customer.Close()
				return nil, err
			}
		}

		customer.Close()

		if err := r.loadOrderItems(order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

const updateOrderQuery = `
UPDATE orders
SET
	status = $1,
	updated_at = $2
WHERE id = $3
`

const restoreStockQuery = `
UPDATE products
SET
	stock = stock + $1,
	updated_at = NOW()
WHERE id = $2
`

func (r *PostgresOrderRepository) Update(
	order *domain.Order,
) error {

	result, err := r.db.Exec(
		updateOrderQuery,
		order.Status,
		order.UpdatedAt,
		order.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrOrderNotFound
	}

	return nil
}

func (r *PostgresOrderRepository) RestoreStock(
	order *domain.Order,
) error {

	for _, item := range order.Items {

		result, err := r.db.Exec(
			restoreStockQuery,
			item.Quantity,
			item.Product.ID,
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
	}

	return nil
}

func (r *PostgresOrderRepository) Delete(
	id uuid.UUID,
) error {
	panic("not implemented")
}
