package cart

import (
	"context"
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

func (s *Service) GetCart(ctx context.Context, userId string) (*cart.Cart, error) {
	userCart, err := s.repo.GetByUserID(ctx, userId)
	if err != nil {
		log.Printf("failed to get cart for user ID %s: %v", userId, err)
		return nil, err
	}
	return userCart, nil
}

func (s *Service) AddProduct(ctx context.Context, userId, sku string) {
	userCart, err := s.repo.GetByUserID(ctx, userId)
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

func (s *Service) RemoveProduct(ctx context.Context, userId, sku string) {
	userCart, err := s.repo.GetByUserID(ctx, userId)
	if err != nil {
		log.Printf("failed to get cart for user ID %s: %v", userId, err)
		return
	}

	userCart.Remove(sku)

	if err := s.repo.Save(ctx, userCart); err != nil {
		log.Printf("failed to save updated cart for user ID %s: %v", userId, err)
	}
}

func (s *Service) CleanCart(ctx context.Context, userId string) (*cart.Cart, error) {
	userCart, err := s.repo.GetByUserID(ctx, userId)
	if err != nil {
		log.Printf("failed to get cart for user ID %s: %v", userId, err)
		return nil, err
	}

	userCart.Clear()

	if err := s.repo.Save(ctx, userCart); err != nil {
		log.Printf("failed to save cleaned cart for user ID %s: %v", userId, err)
		return nil, err
	}

	return userCart, nil
}
