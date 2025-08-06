package cart

type Product struct {
	SKU      string
	Price    int
	Quantity int
}
type Cart struct {
	UserID     string
	TotalPrice int
	Products   []Product
}

// TODO: QuantityOf(sku string) int

// TODO: TotalItems() int

// TODO: IsEmpty() bool

func (c *Cart) Add(p Product) {
	for i := range c.Products {
		if c.Products[i].SKU == p.SKU {
			c.Products[i].Quantity += 1
			c.TotalPrice += p.Price
			return
		}
	}
	p.Quantity = 1
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

func (c *Cart) Total() {
	var total int
	for _, v := range c.Products {
		total += v.Price
	}
	c.TotalPrice = total
}
