package repository

import (
	"context"
	"database/sql"
	"github.com/airsss993/ca-shop-core/internal/domain/cart"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) cart.Repository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) GetByUserID(ctx context.Context, userId string) (*cart.Cart, error) {
	var userCart cart.Cart

	query := `SELECT user_id, total_price FROM carts WHERE user_id=$1`
	row := p.db.QueryRowContext(ctx, query, userId)

	err := row.Scan(&userCart.UserID, &userCart.TotalPrice)
	if err != nil {
		log.Printf("failed to get user_id and total_price for user ID %s: %v", userId, err)
		return nil, err
	}

	query = `SELECT sku, price, quantity FROM cart_items WHERE user_id=$1`
	rows, err := p.db.QueryContext(ctx, query, userId)

	defer rows.Close()

	userCart.Products = make([]cart.Product, 0)

	for rows.Next() {
		var cartProduct cart.Product
		err := rows.Scan(&cartProduct.SKU, &cartProduct.Price, &cartProduct.Quantity)
		if err != nil {
			log.Printf("failed to get products for cart for user ID %s: %v", userId, err)
			return nil, err
		}
		userCart.Products = append(userCart.Products, cartProduct)
	}

	return &userCart, nil
}

func (p *PostgresRepository) Save(ctx context.Context, cart *cart.Cart) error {

	return nil
}
