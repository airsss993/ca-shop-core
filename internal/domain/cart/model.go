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

// TODO: Реализовать метод QuantityOf(sku string) int

// TODO: Реализовать метод Clear()

// TODO: Реализовать метод TotalItems() int

// TODO: Реализовать метод IsEmpty() bool

func (c *Cart) Add(p Product) {
	c.Products = append(c.Products, p)
	c.TotalPrice += p.Price
}

func (c *Cart) Remove(sku string) {
	var newProducts []Product
	for _, p := range c.Products {
		if p.SKU != sku {
			newProducts = append(newProducts, p)
		}
	}
	c.Products = newProducts
}

func (c *Cart) Clear() {
	c.Products = nil
	c.TotalPrice = 0
}

func (c *Cart) Sum() {
	var total int
	for _, v := range c.Products {
		total += v.Price
	}
	c.TotalPrice = total
}
