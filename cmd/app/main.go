package main

import (
	"fmt"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/config"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/database"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/service"
)

func main() {

	productRepository := repository.NewMemoryProductRepository()

	orderRepository := repository.NewMemoryOrderRepository()

	orderService := service.NewOrderService(
		productRepository,
		orderRepository,
	)

	//teste:

	product, err := domain.NewProduct(
		"Micropipeta KASVI",
		"Micropipeta KASVI Multicanal",
		"KASVI",
		3972.79,
		10,
	)
	if err != nil {
		panic(err)
	}

	if err := productRepository.Create(product); err != nil {
		panic(err)
	}

	birthDate := time.Date(
		1995,
		time.March,
		15,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	customer, err := domain.NewCustomer(
		"Danielly Ramos",
		"12345678901",
		birthDate,
		"Rua das Flores, 100",
		"danielly@email.com",
		"83999999999",
	)
	if err != nil {
		panic(err)
	}

	order, err := orderService.CreateOrder(
		customer,
		[]service.CreateOrderItemRequest{
			{
				Product:  product,
				Quantity: 2,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	//exibição de teste:

	fmt.Println("=== LAB SUPPLY API ===")
	fmt.Printf("Cliente: %s\n", customer.Name)
	fmt.Printf("Produto: %s\n", product.Name)
	fmt.Printf("Quantidade comprada: %d\n", order.Items[0].Quantity)
	fmt.Printf("Valor total: R$ %.2f\n", order.Total())
	fmt.Printf("Estoque restante: %d\n", product.Stock)

	//conexão com BD:

	cfg := &config.Config{
		DBHost:     "localhost",
		DBPort:     "5433",
		DBUser:     "postgres",
		DBPassword: "password",
		DBName:     "labsupply",
		DBSSLMode:  "disable",
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("✅ Connected to PostgreSQL!")

}
