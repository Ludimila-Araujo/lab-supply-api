package domain

import "errors"

var (

	// Product erros
	ErrProductNameRequired        = errors.New("product name is required")
	ErrProductDescriptionRequired = errors.New("product description is required")
	ErrProductBrandRequired       = errors.New("product brand is required")
	ErrProductPriceRequired       = errors.New("product price is required, must be greater than zero")
	ErrProductStockRequired       = errors.New("product stock is required, cannot be negative")

	//Customer errors
	ErrCustomerNameRequired      = errors.New("customer name is required")
	ErrCustomerCpfRequired       = errors.New("customer cpf is required")
	ErrCustomerCpfInvalid        = errors.New("customer cpf must contain 11 digits")
	ErrCustomerBirthDateRequired = errors.New("customer birthdate is required")
	ErrCustomerBirthDateInvalid  = errors.New("customer birth date is invalid")
	ErrCustomerUnderAge          = errors.New("customer must be at least 18 years old")
	ErrCustomerOverAge           = errors.New("customer must be no more than 120 years old")
	ErrCustomerAddressRequired   = errors.New("customer address is required")
	ErrCustomerEmailRequired     = errors.New("customer email is required")
	ErrCustomerPhoneRequired     = errors.New("customer phone is required")

	//OredersItem errors:

	ErrOrderItemProductRequired = errors.New("order item product is required")
	ErrOrderItemQuantityInvalid = errors.New("order item product must be greater than zero")

	//Orders erros:

	ErrOrderCustomerRequired = errors.New("order customer is required")
	ErrOrderItemRequired     = errors.New("order item is required")
	ErrOrderCannotBeModified = errors.New("order cannot be modified")
)
