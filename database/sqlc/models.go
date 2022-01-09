// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int32          `json:"id"`
	Name        sql.NullString `json:"name"`
	Price       sql.NullInt64  `json:"price"`
	Description sql.NullString `json:"description"`
	Quantity    sql.NullInt32  `json:"quantity"`
	CreatedAt   time.Time      `json:"created_at"`
}