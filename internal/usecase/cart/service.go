package cart

import (
	"github.com/airsss993/ca-shop-core/internal/domain/cart"
	"log"
)

type Service struct {
	repo    cart.Repository
	catalog cart.ProductCatalog
}

func NewService(repo cart.Repository, catalog cart.ProductCatalog) *Service {
	return &Service{repo: repo, catalog: catalog}
}

func (s *Service) GetCart(userId string) (*cart.Cart, error) {
	userCart, err := s.repo.GetByUserID(userId)
	if err != nil {
		log.Printf("failed to get cart for user ID %s: %v", userId, err)
		return nil, err
	}
	return userCart, nil
}

func (s *Service) AddProduct(userId, sku string) {
	userCart, err := s.repo.GetByUserID(userId)
	if err != nil {
		log.Printf("failed to get cart for user ID %s: %v", userId, err)
		return
	}

	product, err := s.catalog.GetBySKU(sku)
	if err != nil {
		log.Printf("product with sku %s not found: %v", sku, err)
		return
	}

	userCart.Add(product)
}

func (s *Service) RemoveProduct(userId, sku string) {
	userCart, err := s.repo.GetByUserID(userId)
	if err != nil {
		log.Printf("failed to get cart for user ID %s: %v", userId, err)
		return
	}

	userCart.Remove(sku)

	if err := s.repo.Save(userCart); err != nil {
		log.Printf("failed to save updated cart for user ID %s: %v", userId, err)
	}
}

func (s *Service) CleanCart(userId string) (*cart.Cart, error) {
	userCart, err := s.repo.GetByUserID(userId)
	if err != nil {
		log.Printf("failed to get cart for user ID %s: %v", userId, err)
		return nil, err
	}

	userCart.Clear()

	if err := s.repo.Save(userCart); err != nil {
		log.Printf("failed to save cleaned cart for user ID %s: %v", userId, err)
		return nil, err
	}

	return userCart, nil
}
