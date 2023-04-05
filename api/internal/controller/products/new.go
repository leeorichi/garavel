package products

import (
	"context"

	"github.com/leeorichi/getgo/api/internal/model"
	"github.com/leeorichi/getgo/api/internal/repository"
	"github.com/leeorichi/getgo/api/internal/repository/dbmodel"
)

// Controller represents the specification of this pkg
type Controller interface {
	Create(context.Context, CreateInput) (model.Product, error)
	Delete(context.Context, string) (bool, error)
	Show(context.Context, string) (model.Product, error)
	All(context.Context) (dbmodel.ProductSlice, error)
	Update(context.Context, UpdateInput, string) (model.Product, error)
}

// New returns an implementation instance which satisfying Controller
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
