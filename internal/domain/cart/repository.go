package cart

import "context"

type Repository interface {
	GetByUserID(ctx context.Context, userId string) (*Cart, error)
	Save(ctx context.Context, cart *Cart) error
}

type ProductCatalog interface {
	GetBySKU(ctx context.Context, sku string) (Product, error)
}
