package order

import "errors"

var (
	ErrEmptyOrder    = errors.New("empty order")
	ErrInvalidItem   = errors.New("invalid order item")
	ErrPriceMismatch = errors.New("price mismatch")
)
