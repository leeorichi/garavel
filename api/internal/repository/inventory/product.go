package inventory

import (
	"context"
	"fmt"

	"github.com/leeorichi/getgo/api/internal/model"
	"github.com/leeorichi/getgo/api/internal/repository/dbmodel"
	"github.com/leeorichi/getgo/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) CreateProduct(ctx context.Context, m model.Product) (model.Product, error) {
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return model.Product{}, err
	}

	o := dbmodel.Product{
		ID:          id,
		ExternalID:  m.ExternalID,
		Name:        m.Name,
		Description: m.Description,
		Status:      m.Status.String(),
		Price:       m.Price,
	}

	if err := o.Insert(ctx, i.db, boil.Infer()); err != nil {
		return model.Product{}, pkgerrors.WithStack(err)
	}

	m.ID = id
	m.CreatedAt = o.CreatedAt
	m.UpdatedAt = o.UpdatedAt

	return m, nil
}

func (i impl) DeleteProduct(ctx context.Context, eid string) (bool, error) {
	p, err := dbmodel.FindProduct(ctx, i.db, 1)
	if err != nil {
		fmt.Println("===> inventory")
		fmt.Println(err)
		fmt.Println("<=== inventory")

		return false, err
	}
	fmt.Println(p.Name)
	fmt.Println(p.ID)

	p.Delete(ctx, i.db)
	return true, nil
}

// func (i impl) UpdateProduct(ctx context.Context, id int, m model.Product) (model.Product, error) {
// 	// return
// }
