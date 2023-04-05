package inventory

import (
	"context"
	"database/sql"

	"github.com/leeorichi/getgo/api/internal/model"
)

type Repository interface {
	// specification
	CreateProduct(context.Context, model.Product) (model.Product, error)
	DeleteProduct(context.Context, string) (bool, error)
	ShowProduct(context.Context, string) (model.Product, error)
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}

type impl struct {
	db *sql.DB
}
