package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	db "github.com/Sanjungliu/product-api-assesment/database/sqlc"
	"github.com/Sanjungliu/product-api-assesment/internal/httpserver/response"
	"github.com/Sanjungliu/product-api-assesment/internal/product"
)

type Controller struct {
	Product *product.Service
}

func NewController(productService *product.Service) *Controller {
	return &Controller{
		Product: productService,
	}
}

func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request) {
	var param db.AddProductParams

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		response.JSON(w, http.StatusBadRequest, "Bad Request")
	}

	product, err := c.Product.AddProduct(r.Context(), param)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.JSON(w, http.StatusOK, product)
}

func (c *Controller) GetListProduct(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	limit := 5
	offset := 0
	name := ""
	price := ""
	quantity := ""
	createdAt := ""

	if queryValues.Get("limit") != "" {
		limitQuery, err := strconv.Atoi(queryValues.Get("limit"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Bad Request")
			return
		}
		limit = limitQuery
	}
	if queryValues.Get("offset") != "" {
		offsetQuery, err := strconv.Atoi(queryValues.Get("offset"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Bad Request")
			return
		}
		offset = offsetQuery
	}
	if strings.ToUpper(queryValues.Get("name")) == "DESC" {
		name = "name DESC "
	}
	if strings.ToUpper(queryValues.Get("name")) == "ASC" {
		name = "name ASC "
	}
	if strings.ToUpper(queryValues.Get("price")) == "DESC" {
		price = "price DESC "
	}
	if strings.ToUpper(queryValues.Get("price")) == "ASC" {
		price = "price ASC "
	}
	if strings.ToUpper(queryValues.Get("quantity")) == "DESC" {
		quantity = "quantity DESC "
	}
	if strings.ToUpper(queryValues.Get("quantity")) == "ASC" {
		quantity = "quantity ASC "
	}
	if strings.ToUpper(queryValues.Get("createdAt")) == "DESC" {
		createdAt = "created_at DESC "
	}
	if strings.ToUpper(queryValues.Get("createdAt")) == "ASC" {
		createdAt = "created_at ASC "
	}

	products, err := c.Product.GetListProduct(r.Context(), db.ListProductsParams{
		Limit:     limit,
		Offset:    offset,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
		CreatedAt: createdAt,
	})
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.JSON(w, http.StatusOK, products)
}
