package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/controllers/dto"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/service"

	"github.com/google/uuid"
)

//struct:

type ProductController struct {
	productService *service.ProductService
}

//construtor:

func NewProductController(
	productService *service.ProductService,
) *ProductController {

	return &ProductController{
		productService: productService,
	}
}

//CRUD:

func (c *ProductController) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request dto.CreateProductRequest

	// Decodifica o JSON recebido
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Cria a entidade do domínio
	product, err := domain.NewProduct(
		request.Name,
		request.Description,
		request.Brand,
		request.Price,
		request.Stock,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Persiste no banco
	if err := c.productService.Create(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) FindAll(
	w http.ResponseWriter,
	r *http.Request,
) {

	products, err := c.productService.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (c *ProductController) FindByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	idString := r.PathValue("id")

	id, err := uuid.Parse(idString)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := c.productService.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}
