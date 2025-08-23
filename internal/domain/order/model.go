package order

import (
	"errors"
	"time"

	"github.com/airsss993/ca-shop-core/internal/domain/cart"
)

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusValidated OrderStatus = "validated"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusCanceled  OrderStatus = "canceled"
)

type Order struct {
	ID           string
	UserID       string
	Items        []cart.Product
	Price        int64
	Status       OrderStatus
	CreatedAt    time.Time
	ValidatedAt  *time.Time
	PaidAt       *time.Time
	CanceledAt   *time.Time
	CancelReason *string
}

var (
	ErrEmptyOrder        = errors.New("order is empty")
	ErrInvalidItem       = errors.New("order has invalid item")
	ErrDuplicateSKU      = errors.New("order has duplicate SKU")
	ErrPriceMismatch     = errors.New("order price mismatch")
	ErrInvalidTransition = errors.New("invalid status transition")
)

func (o *Order) BasicValidate() error {
	if len(o.Items) == 0 {
		return ErrEmptyOrder
	}

	for _, it := range o.Items {
		if it.SKU == "" || it.Quantity <= 0 || it.Price < 0 {
			return ErrInvalidItem
		}
	}

	if o.TotalFromItems() != o.Price {
		return ErrPriceMismatch
	}

	if o.Status != OrderStatusCreated {
		return ErrInvalidTransition
	}

	return nil
}

func (o *Order) TotalFromItems() int64 {
	var total int64

	for _, it := range o.Items {
		total += it.Price * int64(it.Quantity)
	}

	return total
}
