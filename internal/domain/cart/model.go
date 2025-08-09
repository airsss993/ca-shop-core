package cart

import "time"

type Product struct {
	SKU      string
	Price    int64
	Quantity int
}
type Cart struct {
	UserID     string
	TotalPrice int64
	Products   []Product
	UpdatedAt  time.Time
}

// TODO: QuantityOf(sku string) int

// TODO: TotalItems() int

// TODO: IsEmpty() bool

func (c *Cart) Add(p Product) {
	for i := range c.Products {
		if c.Products[i].SKU == p.SKU {
			c.Products[i].Quantity += 1
			c.RecalculateTotal()
			return
		}
	}
	p.Quantity = 1
	c.Products = append(c.Products, p)
	c.RecalculateTotal()
}

func (c *Cart) Remove(sku string) {
	var newProducts []Product
	for _, p := range c.Products {
		if p.SKU == sku {
			if p.Quantity > 1 {
				p.Quantity = p.Quantity - 1
				newProducts = append(newProducts, p)
			}
			continue
		} else {
			newProducts = append(newProducts, p)
		}
	}
	c.Products = newProducts
	c.RecalculateTotal()
}

func (c *Cart) Clear() {
	c.Products = nil
	c.TotalPrice = 0
}

func (c *Cart) RecalculateTotal() {
	var total int64
	for _, v := range c.Products {
		total += v.Price * int64(v.Quantity)
	}
	c.TotalPrice = total
}
