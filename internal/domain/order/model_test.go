package order_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/airsss993/ca-shop-core/internal/domain/cart"
	"github.com/airsss993/ca-shop-core/internal/domain/order"
)

func TestOrder_BasicValidate(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name    string
		o       order.Order
		wantErr error
	}{
		{
			name:    "empty order",
			o:       order.Order{Status: order.OrderStatusCreated},
			wantErr: order.ErrEmptyOrder,
		},
		{
			name: "invalid item (quantity <= 0)",
			o: order.Order{
				Status: order.OrderStatusCreated,
				Items: []cart.Product{
					{SKU: "A", Quantity: 0, Price: 100},
				},
				Price: 100,
			},
			wantErr: order.ErrInvalidItem,
		},
		{
			name: "price mismatch",
			o: order.Order{
				Status: order.OrderStatusCreated,
				Items: []cart.Product{
					{SKU: "A", Quantity: 1, Price: 100},
					{SKU: "B", Quantity: 2, Price: 150},
				},
				Price: 100,
			},
			wantErr: order.ErrPriceMismatch,
		},
		{
			name: "valid order passes",
			o: order.Order{
				Status:    order.OrderStatusCreated,
				CreatedAt: now,
				Items: []cart.Product{
					{SKU: "A", Quantity: 1, Price: 100},
					{SKU: "B", Quantity: 2, Price: 150},
				},
				Price: 400,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.o.BasicValidate()
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
