// source: product.sql

package db

import (
	"context"
)

type AddProductParams struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (q *Queries) AddProduct(ctx context.Context, params AddProductParams) (Product, error) {
	query := `
	INSERT INTO product (
		name,
		price,
		description,
		quantity
	) VALUES (
		$1, $2, $3, $4
	) RETURNING id, name, price, description, quantity, created_at
	`
	row := q.db.QueryRowContext(ctx, query,
		params.Name,
		params.Price,
		params.Description,
		params.Quantity,
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

type ListProductsParams struct {
	Limit     int
	Offset    int
	Price     string
	Name      string
	Quantity  string
	CreatedAt string
}

func (q *Queries) ListProducts(ctx context.Context, params ListProductsParams) ([]Product, error) {
	counter := 0
	query := `
		SELECT id, name, price, description, quantity, created_at FROM product
	`

	orderBy := ` ORDER BY `
	if params.Name != "" {
		orderBy += params.Name
		counter++
	} else if params.Price != "" {
		orderBy += params.Price
		counter++
	} else if params.Quantity != "" {
		orderBy += params.Quantity
		counter++
	}

	if counter != 0 {
		query += orderBy
	}

	query += `LIMIT $1 OFFSET $2`

	rows, err := q.db.QueryContext(ctx, query, params.Limit, params.Offset)
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
