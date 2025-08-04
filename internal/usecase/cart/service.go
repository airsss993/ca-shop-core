package cart

import (
	"github.com/airsss993/ca-shop-core/internal/domain/cart"
	"log"
)

type Service struct {
	repo cart.CartRepository
}

func NewService(repo cart.CartRepository) *Service {
	return &Service{repo: repo}
}

// TODO: AddProduct(userID, sku string)

func (s *Service) RemoveProduct(userId, sku string) {
	cart, err := s.repo.GetByUserID(userId)
	if err != nil {
		log.Printf("failed to get cart for user ID %s: %v", userId, err)
		return
	}

	cart.Remove(sku)

	if err := s.repo.Save(cart); err != nil {
		log.Printf("failed to save updated cart for user ID %s: %v", userId, err)
	}
}

// TODO: GetCart(userID string)
// TODO: ClearCart(userID string)
