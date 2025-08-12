package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/airsss993/ca-shop-core/internal/domain/cart"
	"github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) GetCartByUserID(ctx context.Context, userId string) (*cart.Cart, error) {
	var userCart cart.Cart

	query := `SELECT user_id, total_price FROM carts WHERE user_id=$1`
	row := p.db.QueryRowContext(ctx, query, userId)

	err := row.Scan(&userCart.UserID, &userCart.TotalPrice)
	if errors.Is(err, sql.ErrNoRows) {
		return &cart.Cart{
			UserID:     userId,
			TotalPrice: 0,
			Products:   []cart.Product{},
			UpdatedAt:  time.Now(),
		}, nil
	}
	if err != nil {
		return nil, err
	}

	query = `SELECT sku, price, quantity FROM cart_items WHERE user_id=$1`
	rows, err := p.db.QueryContext(ctx, query, userId)

	userCart.Products = make([]cart.Product, 0)

	for rows.Next() {
		var cartProduct cart.Product
		err := rows.Scan(&cartProduct.SKU, &cartProduct.Price, &cartProduct.Quantity)
		if err != nil {
			return nil, err
		}
		userCart.Products = append(userCart.Products, cartProduct)
	}

	defer rows.Close()

	userCart.RecalculateTotal()

	return &userCart, nil
}

func (p *PostgresRepository) SaveCart(ctx context.Context, cart *cart.Cart) error {
	cart.RecalculateTotal()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        INSERT INTO carts (user_id, total_price, updated_at)
        VALUES ($1, $2, $3)
        ON CONFLICT (user_id)
        DO UPDATE SET total_price=EXCLUDED.total_price, updated_at=EXCLUDED.updated_at;
    `
	if _, err = tx.ExecContext(ctx, query, cart.UserID, cart.TotalPrice, cart.UpdatedAt); err != nil {
		return fmt.Errorf("upsert cart header: %w", err)
	}

	for _, v := range cart.Products {
		query = `
            INSERT INTO cart_items (user_id, sku, price, quantity)
            VALUES ($1, $2, $3, $4)
            ON CONFLICT (user_id, sku)
            DO UPDATE SET price=EXCLUDED.price, quantity=EXCLUDED.quantity;
        `
		if _, err = tx.ExecContext(ctx, query, cart.UserID, v.SKU, v.Price, v.Quantity); err != nil {
			return fmt.Errorf("upsert cart item sku=%s: %w", v.SKU, err)
		}
	}

	actualSku := make([]string, 0, len(cart.Products))
	for _, v := range cart.Products {
		actualSku = append(actualSku, v.SKU)
	}

	if len(actualSku) == 0 {
		query = `DELETE FROM cart_items WHERE user_id=$1;`
		if _, err = tx.ExecContext(ctx, query, cart.UserID); err != nil {
			return fmt.Errorf("delete all items for user_id=%s: %w", cart.UserID, err)
		}
	} else {
		query = `DELETE FROM cart_items WHERE user_id=$1 AND sku <> ALL($2);`
		if _, err = tx.ExecContext(ctx, query, cart.UserID, pq.Array(actualSku)); err != nil {
			return fmt.Errorf("delete missing items for user_id=%s: %w", cart.UserID, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
