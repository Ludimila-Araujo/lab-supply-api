package seed

import "database/sql"

func Run(db *sql.DB) error {
	if err := SeedProducts(db); err != nil {
		return err
	}

	if err := SeedCustomers(db); err != nil {
		return err
	}

	if err := SeedOrders(db); err != nil {
		return err
	}

	return nil
}
