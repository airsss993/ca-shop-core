package order

import (
	"time"

	"github.com/airsss993/ca-shop-core/internal/domain/cart"
)

type Status string

const (
	StatusCreated   Status = "created"
	StatusValidated Status = "validated"
	StatusPaid      Status = "paid"
	StatusCanceled  Status = "canceled"
)

type Order struct {
	ID           string
	UserID       string
	Items        []cart.Product
	Price        int64
	Status       Status
	CreatedAt    time.Time
	ValidatedAt  *time.Time
	PaidAt       *time.Time
	CanceledAt   *time.Time
	CancelReason *string
}

func (o *Order) TotalFromItems() int64 {
	var total int64

	for _, it := range o.Items {
		total += it.Price * int64(it.Quantity)
	}

	return total
}

func (o *Order) MarkAsValidated() {
	o.Status = StatusValidated
	now := time.Now()
	o.ValidatedAt = &now
}

func (o *Order) MarkAsPaid() {
	o.Status = StatusPaid
	now := time.Now()
	o.PaidAt = &now
}

func (o *Order) Cancel(reason string) {
	o.Status = StatusCanceled
	now := time.Now()
	o.CanceledAt = &now
	o.CancelReason = &reason
}
