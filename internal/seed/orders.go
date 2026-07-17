package seed

import (
	"database/sql"
	"fmt"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
)

type seedOrder struct {
	CustomerCPF string
	Status      string
	Items       []seedOrderItem
}

type seedOrderItem struct {
	ProductName string
	Quantity    int
}

var orders = []seedOrder{
	{
		CustomerCPF: "52998224725",
		Status:      "PENDING",
		Items: []seedOrderItem{
			{
				ProductName: "Micropipeta Monocanal 100 µL",
				Quantity:    1,
			},
			{
				ProductName: "Ponteira 200 µL",
				Quantity:    3,
			},
		},
	},

	{
		CustomerCPF: "12345678909",
		Status:      "PAID",
		Items: []seedOrderItem{
			{
				ProductName: "Agarose Molecular Biology",
				Quantity:    2,
			},
			{
				ProductName: "Master Mix PCR 2X",
				Quantity:    1,
			},
		},
	},

	{
		CustomerCPF: "11144477735",
		Status:      "CANCELED",
		Items: []seedOrderItem{
			{
				ProductName: "Kit ELISA IgG",
				Quantity:    1,
			},
		},
	},
}

func SeedOrders(db *sql.DB) error {

	customerRepository := repository.NewPostgresCustomerRepository(db)
	productRepository := repository.NewPostgresProductRepository(db)
	orderRepository := repository.NewPostgresOrderRepository(db)

	created := 0

	for _, o := range orders {

		customer, err := customerRepository.FindByCPF(o.CustomerCPF)
		if err != nil {
			return err
		}

		order, err := domain.NewOrder(customer)
		if err != nil {
			return err
		}

		for _, item := range o.Items {

			product, err := productRepository.FindByName(item.ProductName)
			if err != nil {
				return err
			}

			orderItem, err := domain.NewOrderItem(
				product,
				item.Quantity,
			)
			if err != nil {
				return err
			}

			if err := order.AddItem(orderItem); err != nil {
				return err
			}
		}

		switch o.Status {

		case "PAID":
			if err := order.Pay(); err != nil {
				return err
			}

		case "CANCELED":
			if err := order.Cancel(); err != nil {
				return err
			}
		}

		if err := orderRepository.Create(order); err != nil {
			return err
		}

		created++
	}

	fmt.Printf(
		"Orders: %d created\n",
		created,
	)

	return nil
}
