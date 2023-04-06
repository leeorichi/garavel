package products

import (
	"context"

	"github.com/leeorichi/getgo/api/internal/model"
)

// find prod by external id
func (i impl) Show(ctx context.Context, eid string) (model.Product, error) {
	return i.repo.Inventory().ShowProduct(ctx, eid)
}
