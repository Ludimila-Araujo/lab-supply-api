package domain

import "errors"

var (

	// =========================================================================
	// Product validation
	// =========================================================================

	ErrProductNameRequired        = errors.New("product name is required")
	ErrProductDescriptionRequired = errors.New("product description is required")
	ErrProductBrandRequired       = errors.New("product brand is required")
	ErrProductPriceRequired       = errors.New("product price must be greater than zero")
	ErrProductStockRequired       = errors.New("product stock cannot be negative")

	// =========================================================================
	// Customer validation
	// =========================================================================

	ErrCustomerNameRequired         = errors.New("customer name is required")
	ErrCustomerCpfRequired          = errors.New("customer cpf is required")
	ErrCustomerCpfInvalid           = errors.New("customer cpf must contain 11 digits")
	ErrCustomerBirthDateRequired    = errors.New("customer birth date is required")
	ErrCustomerBirthDateInvalid     = errors.New("customer birth date is invalid")
	ErrCustomerUnderAge             = errors.New("customer must be at least 18 years old")
	ErrCustomerOverAge              = errors.New("customer must be no more than 120 years old")
	ErrCustomerAddressRequired      = errors.New("customer address is required")
	ErrCustomerEmailRequired        = errors.New("customer email is required")
	ErrCustomerPhoneRequired        = errors.New("customer phone is required")
	ErrCustomerPasswordHashRequired = errors.New("customer password hash is required")

	// =========================================================================
	// Order validation
	// =========================================================================

	ErrOrderCustomerRequired = errors.New("order customer is required")
	ErrOrderItemRequired     = errors.New("order item is required")
	ErrOrderCannotBeModified = errors.New("order cannot be modified")
	ErrOrderInvalidStatus    = errors.New("invalid order status")

	// =========================================================================
	// Order Item validation
	// =========================================================================

	ErrOrderItemProductRequired = errors.New("order item product is required")
	ErrOrderItemQuantityInvalid = errors.New("order item quantity must be greater than zero")

	// =========================================================================
	// Business rules
	// =========================================================================

	// Customer
	ErrCustomerAlreadyExists = errors.New("customer already exists")
	ErrCustomerNotFound      = errors.New("customer not found")

	// Product
	ErrProductAlreadyExists     = errors.New("product already exists")
	ErrProductNotFound          = errors.New("product not found")
	ErrProductInsufficientStock = errors.New("insufficient product stock")

	// Order
	ErrOrderAlreadyExists = errors.New("order already exists")
	ErrOrderNotFound      = errors.New("order not found")
)
