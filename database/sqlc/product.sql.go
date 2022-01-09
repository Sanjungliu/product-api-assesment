// source: product.sql

package db

import (
	"context"
)

const addProduct = `-- name: AddProduct :one
INSERT INTO product (
    name,
	price,
	description,
	quantity
) VALUES (
    $1, $2, $3, $4
) RETURNING id, name, price, description, quantity, created_at
`

type AddProductParams struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (q *Queries) AddProduct(ctx context.Context, arg AddProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, addProduct,
		arg.Name,
		arg.Price,
		arg.Description,
		arg.Quantity,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Description,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, price, description, quantity, created_at FROM product
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Description,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, price, description, quantity, created_at FROM product
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Description,
			&i.Quantity,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
