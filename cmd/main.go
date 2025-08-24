package main

import (
	"log"

	"github.com/airsss993/ca-shop-core/internal/adapters/cart/http"
	"github.com/airsss993/ca-shop-core/internal/adapters/cart/repository"
	db2 "github.com/airsss993/ca-shop-core/internal/infra/db"
	"github.com/airsss993/ca-shop-core/internal/usecase/cart"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := db2.InitDB()
	defer db.Close()

	// Cart Module
	cartRepo := repository.NewCartRepository(db)
	cartService := cart.NewService(cartRepo, nil)
	cartHandler := http.NewCartHandler(*cartService)
	cartHandler.RegisterRoutes(&router.RouterGroup)

	// Order Module

	err := router.Run()
	if err != nil {
		log.Fatal("failed to start HTTP server")
		return
	}
}
