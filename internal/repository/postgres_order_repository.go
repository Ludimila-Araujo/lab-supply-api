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

func (r *PostgresOrderRepository) FindByID(
	id uuid.UUID,
) (*domain.Order, error) {
	panic("not implemented")
}

func (r *PostgresOrderRepository) FindAll(
	limit, offset int,
) ([]*domain.Order, error) {
	panic("not implemented")
}

func (r *PostgresOrderRepository) Update(
	order *domain.Order,
) error {
	panic("not implemented")
}

func (r *PostgresOrderRepository) Delete(
	id uuid.UUID,
) error {
	panic("not implemented")
}
