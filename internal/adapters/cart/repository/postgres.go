package repository

import (
	"database/sql"
	"github.com/airsss993/ca-shop-core/internal/domain/cart"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) cart.Repository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) GetByUserID(userId string) (*cart.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresRepository) Save(cart *cart.Cart) error {
	//TODO implement me
	panic("implement me")
}
