package products

import (
	"context"
)

// Create creates new product
func (i impl) Delete(ctx context.Context, eid string) (bool, error) {
	return i.repo.Inventory().DeleteProduct(ctx, eid)
}
