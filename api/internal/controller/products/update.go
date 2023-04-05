package products

import (
	"context"

	"github.com/leeorichi/getgo/api/internal/model"
)

// CreateInput represents for input to create Product
type UpdateInput struct {
	Name        string
	Description string
	Price       int64
}

// Create creates new product
func (i impl) Update(ctx context.Context, input UpdateInput, eid string) (model.Product, error) {
	return i.repo.Inventory().UpdateProduct(ctx, model.Product{
		Price:       input.Price,
		Description: input.Description,
		Name:        input.Name,
	}, eid)
}
