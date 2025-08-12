package http

import (
	"context"
	"log"
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

	log.Println("USER ID -", userId)

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
