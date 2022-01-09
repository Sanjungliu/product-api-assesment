package product

import (
	"context"
	"database/sql"

	db "github.com/Sanjungliu/product-api-assesment/database/sqlc"
)

func NewService(storage *sql.DB) *Service {
	return &Service{
		storage: *db.New(storage),
	}
}

type Service struct {
	storage db.Queries
}

func (s *Service) AddProduct(ctx context.Context, arg db.AddProductParams) (*db.Product, error) {
	product, err := s.storage.AddProduct(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *Service) GetListProduct(ctx context.Context, arg db.ListProductsParams) (*[]db.Product, error) {
	products, err := s.storage.ListProducts(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &products, nil
}
