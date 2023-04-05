package inventory

import (
	"context"
	"fmt"

	"github.com/leeorichi/getgo/api/internal/model"
	"github.com/leeorichi/getgo/api/internal/repository/dbmodel"
	"github.com/leeorichi/getgo/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

	p, err := dbmodel.Products(qm.Where("external_id = ?", eid)).One(ctx, i.db)
	// p, err := dbmodel.FindProduct(ctx, i.db, 1)
	if err != nil {
		fmt.Println("===> inventory")
		fmt.Println(err)
		fmt.Println("<=== inventory")

		return false, err
	}

	p.Delete(ctx, i.db)
	return true, nil
}

// ShowProduct
func (i impl) ShowProduct(ctx context.Context, eid string) (model.Product, error) {

	prod := model.Product{}

	p, err := dbmodel.Products(qm.Where("external_id = ?", eid)).One(ctx, i.db)

	if err != nil {
		fmt.Println("===> inventory")
		fmt.Println(err)
		fmt.Println("<=== inventory")
		return prod, err
	}

	prod = model.Product{
		ExternalID:  p.ExternalID,
		Name:        p.Name,
		Price:       p.Price,
		Status:      model.ProductStatus(p.Status),
		Description: p.Description,
	}
	return prod, nil
}

func (i impl) AllProduct(ctx context.Context) (dbmodel.ProductSlice, error) {
	if p, err := dbmodel.Products().All(ctx, i.db); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (i impl) UpdateProduct(ctx context.Context, m model.Product, eid string) (model.Product, error) {
	p, err := dbmodel.Products(qm.Where("external_id = ?", eid)).One(ctx, i.db)
	if err != nil {
		return model.Product{}, err
	}

	p.Name = m.Name
	p.Description = m.Description
	p.Price = m.Price

	if _, err := p.Update(ctx, i.db, boil.Infer()); err != nil {
		return model.Product{}, err
	}
	return m, nil
}
