package db

import (
	"context"
	"testing"

	"github.com/Sanjungliu/product-api-assesment/pkg"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	arg := AddProductParams{
		Name:        pkg.RandomString(8),
		Price:       int(pkg.RandomInt(10000, 100000)),
		Description: "Terbuat dari bahan pilihan",
		Quantity:    int(pkg.RandomInt(1, 10)),
	}

	product, err := testQueries.AddProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.Quantity, product.Quantity)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func Test_AddProduct(t *testing.T) {
	createRandomProduct(t)
}

func Test_ListProducts(t *testing.T) {
	count := 5
	for i := 0; i < count; i++ {
		createRandomProduct(t)
	}

	args := ListProductsParams{
		Limit:  5,
		Offset: 0,
	}

	products, err := testQueries.ListProducts(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, products)
	require.Equal(t, count, len(products))

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}
