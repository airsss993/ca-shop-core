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
	pgRepo := repository.NewPostgresRepository(db)
	// todo: сделать каталог
	cartService := cart.NewService(pgRepo, nil)
	cartHandler := http.NewCartHandler(*cartService)
	cartHandler.RegisterRoutes(&router.RouterGroup)

	err := router.Run()
	if err != nil {
		log.Fatal("failed to start HTTP server")
		return
	}
}
