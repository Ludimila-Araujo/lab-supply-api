package dto

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
