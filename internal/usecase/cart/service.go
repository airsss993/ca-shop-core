package cart

import (
	"context"

	"github.com/airsss993/ca-shop-core/internal/domain/cart"
)

type Service struct {
	repo    cart.Repository
	catalog cart.ProductCatalog
}

func NewService(repo cart.Repository, catalog cart.ProductCatalog) *Service {
	return &Service{repo: repo, catalog: catalog}
}

func (s *Service) GetCart(ctx context.Context, userId string) (*cart.Cart, error) {
	userCart, err := s.repo.GetCartByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return userCart, nil
}

func (s *Service) AddProduct(ctx context.Context, userId, sku string) error {
	userCart, err := s.repo.GetCartByUserID(ctx, userId)
	if err != nil {
		return err
	}

	product, err := s.catalog.GetBySKU(ctx, sku)
	if err != nil {
		return err
	}

	userCart.Add(product)

	err = s.repo.SaveCart(ctx, userCart)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveProduct(ctx context.Context, userId, sku string) error {
	userCart, err := s.repo.GetCartByUserID(ctx, userId)
	if err != nil {
		return err
	}

	userCart.Remove(sku)

	if err := s.repo.SaveCart(ctx, userCart); err != nil {
		return err
	}

	return nil
}

func (s *Service) CleanCart(ctx context.Context, userId string) (*cart.Cart, error) {
	userCart, err := s.repo.GetCartByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	userCart.Clear()

	if err := s.repo.SaveCart(ctx, userCart); err != nil {
		return nil, err
	}

	return userCart, nil
}
