package main

import (
	"log"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/config"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/controllers"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/database"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/routes"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/server"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/service"
)

func main() {

	// Configuração do banco
	cfg := config.Load()

	// Conexão
	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Para produtos:

	// Repositório PostgreSQL
	productRepository := repository.NewPostgresProductRepository(db)

	// Service:
	productService := service.NewProductService(productRepository)

	//Controller:
	productController := controllers.NewProductController(productService)

	//para clientes:

	// Repositório
	customerRepository := repository.NewPostgresCustomerRepository(db)

	// Service
	customerService := service.NewCustomerService(customerRepository)

	// Controller
	customerController := controllers.NewCustomerController(customerService)

	//Server:

	srv := server.NewServer()

	// Routes
	routes.RegisterProductRoutes(
		srv.Mux(),
		productController,
	)

	routes.RegisterCustomerRoutes(
		srv.Mux(),
		customerController,
	)

	log.Println("🚀 Lab Supply API running on :8080")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
