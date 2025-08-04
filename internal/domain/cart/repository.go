package cart

type CartRepository interface {
	GetByUserID(userId string) (*Cart, error)
	Save(cart *Cart) error
}
