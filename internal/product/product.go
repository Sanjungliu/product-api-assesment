package product

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/Sanjungliu/product-api-assesment/database/sqlc"
	"github.com/go-redis/redis"
)

func NewService(storage *sql.DB, redis *redis.Client) *Service {
	return &Service{
		storage: *db.New(storage),
		redis:   redis,
	}
}

type Service struct {
	storage db.Queries
	redis   *redis.Client
}

func (s *Service) AddProduct(ctx context.Context, params db.AddProductParams) (*db.Product, error) {
	product, err := s.storage.AddProduct(ctx, params)
	if err != nil {
		return nil, err
	}

	s.redis.FlushAll()
	return &product, nil
}

func (s *Service) GetListProduct(ctx context.Context, params db.ListProductsParams) (*[]db.Product, error) {
	key := fmt.Sprintf("%d.%d.%s.%s.%s.%s", params.Limit, params.Offset, params.Price, params.Name, params.Quantity, params.CreatedAt)

	cache, errCache := s.redis.Get(key).Result()
	if errCache != nil {
		products, err := s.storage.ListProducts(ctx, params)
		if err != nil {
			return nil, err
		}

		json, errMarshal := json.Marshal(products)
		if errMarshal != nil {
			return nil, errMarshal
		}

		if errCache := s.redis.Set(key, json, 30*time.Minute).Err(); errCache != nil {
			return nil, errCache
		}

		return &products, nil
	}

	var products []db.Product

	if errUnmarshal := json.Unmarshal([]byte(cache), &products); errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &products, nil
}
