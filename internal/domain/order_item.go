package domain

// representação de item dentro pedido realizado por cliente

type OrderItem struct {
	Product   *Product
	Quantity  int
	UnitPrice float64
}

func NewOrderItem(
	product *Product,
	quantity int,
) (*OrderItem, error) {

	if product == nil {
		return nil, ErrOrderItemProductRequired
	}

	if quantity <= 0 {
		return nil, ErrOrderItemQuantityInvalid
	}

	return &OrderItem{
		Product:   product,
		Quantity:  quantity,
		UnitPrice: product.Price,
	}, nil
}
