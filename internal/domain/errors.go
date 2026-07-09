package domain

import "errors"

var (
	ErrProductNameRequired        = errors.New("product name is required")
	ErrProductDescriptionRequired = errors.New("product description is required")
	ErrProductBrandRequired       = errors.New("prduct brand is required")
	ErrProductPriceRequired       = errors.New("product price is required, must be greater than zero")
	ErrProductStockRequired       = errors.New("prodcut stock is required, cannot be negative")
)
