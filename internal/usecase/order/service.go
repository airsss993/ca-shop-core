package order

import (
	"context"
	"time"

	"github.com/airsss993/ca-shop-core/internal/domain/order"
)

type Service struct {
	repo order.Repository
}

func NewService(repo order.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Validate(o *order.Order) error {
	if len(o.Items) == 0 {
		return order.ErrEmptyOrder
	}

	for _, item := range o.Items {
		if item.SKU == "" || item.Quantity <= 0 || item.Price < 0 {
			return order.ErrInvalidItem
		}
	}

	if o.TotalFromItems() != o.Price {
		return order.ErrPriceMismatch
	}

	return nil
}

func (s *Service) Create(ctx context.Context, o *order.Order) error {
	if err := s.Validate(o); err != nil {
		return err
	}

	o.Status = order.StatusCreated
	o.CreatedAt = time.Now()

	if err := s.repo.Create(ctx, o); err != nil {
		return err
	}

	return nil
}

func (s *Service) FindByID(ctx context.Context, id string) (*order.Order, error) {
	o, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (s *Service) FindByUserID(ctx context.Context, userID string) ([]order.Order, error) {
	return s.repo.GetByUserID(ctx, userID)
}

func (s *Service) ChangeStatus(ctx context.Context, id string, status order.Status) error {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *Service) Cancel(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
