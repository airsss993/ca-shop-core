package http

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/airsss993/ca-shop-core/internal/adapters/cart/http/dto"
	uc "github.com/airsss993/ca-shop-core/internal/usecase/cart"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service uc.Service
}

const defaultCartRead = 200 * time.Millisecond

func NewCartHandler(service uc.Service) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/cart/:userid", h.GetCart)
	rg.POST("/cart/:userid/items/:sku", h.AddToCart)
	rg.DELETE("/cart/:userid/items/:sku", h.DeleteFromCart)
	rg.DELETE("/cart/:userid", h.ClearCart)
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userId := c.Param("userid")
	if strings.TrimSpace(userId) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user_id",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), defaultCartRead)
	defer cancel()

	userCart, err := h.service.GetCart(ctx, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := dto.ToCartResponse(*userCart)

	c.JSON(http.StatusOK, gin.H{
		"cart": resp,
	})
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userId := c.Param("userid")
	if strings.TrimSpace(userId) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user_id",
		})
		return
	}

	itemSku := c.Param("sku")
	if strings.TrimSpace(itemSku) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid sku",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), defaultCartRead)
	defer cancel()

	err := h.service.AddProduct(ctx, userId, itemSku)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to add item",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "item successfully added to cart",
	})
}

// TODO: DELETE /cart/:userid/items/:sku — убрать 1 шт

func (h *CartHandler) DeleteFromCart(c *gin.Context) {
	userId := c.Param("userid")
	if strings.TrimSpace(userId) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user_id",
		})
		return
	}

	itemSku := c.Param("sku")
	if strings.TrimSpace(itemSku) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid sku",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), defaultCartRead)
	defer cancel()

	err := h.service.RemoveProduct(ctx, userId, itemSku)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to remove item",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully removed from cart",
	})
}

// TODO: DELETE /cart/:userid — очистить корзину

func (h *CartHandler) ClearCart(c *gin.Context) {
	userId := c.Param("userid")
	if strings.TrimSpace(userId) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user_id",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), defaultCartRead)
	defer cancel()

	err := h.service.ClearCart(ctx, userId)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to clear cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully cleared cart",
	})
}
