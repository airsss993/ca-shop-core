package repository

import (
	"context"
	"database/sql"

	"github.com/airsss993/ca-shop-core/internal/domain/order"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) Create(ctx context.Context, o *order.Order) error {
	return nil
}

func (p *PostgresRepository) GetByID(ctx context.Context, id string) (*order.Order, error) {
	return nil, nil
}

func (p *PostgresRepository) GetByUserID(ctx context.Context, userID string) ([]order.Order, error) {
	return nil, nil
}

func (p *PostgresRepository) UpdateStatus(ctx context.Context, id string, status order.Status) error {
	return nil
}

func (p *PostgresRepository) Delete(ctx context.Context, id string) error {
	return nil
}
