package cart

import "context"

type Repository interface {
	GetCartByUserID(ctx context.Context, userId string) (*Cart, error)
	SaveCart(ctx context.Context, cart *Cart) error
}

type ProductCatalog interface {
	GetBySKU(ctx context.Context, sku string) (Product, error)
}
