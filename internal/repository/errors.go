package repository

import (
	"errors"
)

var (

	//produtos
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrProductNotFound      = errors.New("product not found")

	//pedido

	ErrOrderAlreadyExists = errors.New("order already exists")
	ErrOrderNotFound      = errors.New("order not found")

	//cliente:

	ErrCustomerNotFound = errors.New("customer not found")
)
