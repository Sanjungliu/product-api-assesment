package main

import (
	"database/sql"

	"github.com/Sanjungliu/product-api-assesment/internal/app"
	"github.com/Sanjungliu/product-api-assesment/internal/product"
)

func buildInternalService(db *sql.DB) *app.App {
	productService := product.NewService(db)

	return &app.App{
		Product: *productService,
	}
}
