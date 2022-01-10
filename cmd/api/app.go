package main

import (
	"database/sql"

	"github.com/Sanjungliu/product-api-assesment/internal/app"
	"github.com/Sanjungliu/product-api-assesment/internal/product"
	"github.com/go-redis/redis"
)

func buildInternalService(db *sql.DB, redis *redis.Client) *app.App {
	productService := product.NewService(db, redis)

	return &app.App{
		Product: *productService,
	}
}
