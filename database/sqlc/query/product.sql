-- name: AddProduct :one
INSERT INTO product (
    name,
	price,
	description,
	quantity
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM product
ORDER BY id
LIMIT $1
OFFSET $2;