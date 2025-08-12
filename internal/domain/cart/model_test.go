package cart

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCart_Add(t *testing.T) {
	c := Cart{
		UserID:     "user124",
		TotalPrice: 0,
		Products:   []Product{},
		UpdatedAt:  time.Now(),
	}

	p := Product{
		SKU:   "sku1",
		Price: 100,
	}

	c.Add(p)

	assert.Len(t, c.Products, 1)
	assert.Equal(t, 1, c.Products[0].Quantity)
	assert.Equal(t, int64(100), c.TotalPrice)

	c.Add(p)

	assert.Len(t, c.Products, 1)
	assert.Equal(t, 2, c.Products[0].Quantity)
	assert.Equal(t, int64(200), c.TotalPrice)
}

func TestCart_Remove(t *testing.T) {
	c := Cart{UserID: "user124"}
	p := Product{
		SKU:   "sku1",
		Price: 100,
	}
	c.Add(p)
	c.Add(p)

	c.Remove("sku1")
	assert.Len(t, c.Products, 1)
	assert.Equal(t, 1, c.Products[0].Quantity)
	assert.Equal(t, int64(100), c.TotalPrice)

	c.Remove("sku1")
	assert.Len(t, c.Products, 0)
	assert.Equal(t, int64(0), c.TotalPrice)
}

func TestCart_Clear(t *testing.T) {
	c := Cart{UserID: "user124"}
	p := Product{
		SKU:   "sku1",
		Price: 100,
	}
	c.Add(p)

	c.Clear()
	assert.Len(t, c.Products, 0)
	assert.Equal(t, int64(0), c.TotalPrice)
	now := time.Now()
	assert.WithinDuration(t, now, c.UpdatedAt, time.Second)
}

func TestCart_RecalculateTotal(t *testing.T) {
	c := Cart{UserID: "user124"}
	p := Product{
		SKU:   "sku1",
		Price: 100,
	}

	c.Add(p)
	p = Product{
		SKU:   "sku2",
		Price: 100,
	}
	c.Add(p)

	var total int64

	for _, v := range c.Products {
		total += v.Price * int64(v.Quantity)
	}

	c.RecalculateTotal()

	assert.Equal(t, total, c.TotalPrice)

}
