package dto

import (
	"time"

	"github.com/airsss993/ca-shop-core/internal/domain/cart"
)

type ProductDTO struct {
	SKU      string `json:"sku"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

type CartDTOResp struct {
	UserID     string       `json:"user_id"`
	TotalPrice int64        `json:"total_price"`
	Products   []ProductDTO `json:"products"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

func ToCartResponse(c cart.Cart) CartDTOResp {
	products := make([]ProductDTO, 0, len(c.Products))
	for _, p := range c.Products {
		products = append(products, ProductDTO{
			SKU:      p.SKU,
			Price:    p.Price,
			Quantity: p.Quantity,
		})
	}
	return CartDTOResp{
		UserID:     c.UserID,
		TotalPrice: c.TotalPrice,
		Products:   products,
		UpdatedAt:  c.UpdatedAt,
	}
}
