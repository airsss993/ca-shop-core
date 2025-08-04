package cart

type Product struct {
	SKU   string
	Price int
}
type Cart struct {
	UserID     string
	TotalPrice int
	Products   []Product
}

// TODO: Реализовать метод HasProduct(sku string) bool

// TODO: Реализовать метод QuantityOf(sku string) int

// TODO: Реализовать метод Clear()

// TODO: Реализовать метод TotalItems() int

// TODO: Реализовать метод IsEmpty() bool

func (c *Cart) Remove(SKU string) {
	var newProducts []Product
	for _, p := range c.Products {
		if p.SKU != SKU {
			newProducts = append(newProducts, p)
		}
	}
	c.Products = newProducts
}

func (c *Cart) Sum() {
	var total int
	for _, v := range c.Products {
		total += v.Price
	}
	c.TotalPrice = total
}
