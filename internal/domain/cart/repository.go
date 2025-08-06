package cart

type Repository interface {
	GetByUserID(userId string) (*Cart, error)
	Save(cart *Cart) error
}

type ProductCatalog interface {
	GetBySKU(sku string) (Product, error)
}
