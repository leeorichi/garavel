package products

import (
	"context"

	"github.com/leeorichi/getgo/api/internal/repository/dbmodel"
)

// find prod by external id
func (i impl) All(ctx context.Context) (dbmodel.ProductSlice, error) {
	return i.repo.Inventory().AllProduct(ctx)
}
