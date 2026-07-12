package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/controllers/dto"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/service"

	"github.com/google/uuid"
)

type CustomerController struct {
	customerService *service.CustomerService
}

func NewCustomerController(
	customerService *service.CustomerService,
) *CustomerController {

	return &CustomerController{
		customerService: customerService,
	}
}

func (c *CustomerController) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request dto.CreateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	birthDate, err := time.Parse(
		"2006-01-02",
		request.BirthDate,
	)

	if err != nil {
		http.Error(w, "invalid birth date", http.StatusBadRequest)
		return
	}

	customer, err := c.customerService.Create(
		request.Name,
		request.CPF,
		birthDate,
		request.Address,
		request.Email,
		request.Phone,
		request.Password,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (c *CustomerController) FindAll(
	w http.ResponseWriter,
	r *http.Request,
) {

	customers, err := c.customerService.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (c *CustomerController) FindByID(
	w http.ResponseWriter,
	r *http.Request,
) {

	idString := r.PathValue("id")

	id, err := uuid.Parse(idString)
	if err != nil {
		http.Error(w, "invalid customer id", http.StatusBadRequest)
		return
	}

	customer, err := c.customerService.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
