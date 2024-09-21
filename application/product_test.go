package application_test

import (
	"testing"

	"github.com/emiliosheinz/exagonal-architecture/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)

	product.Status = "invalid"
	isValid, err = product.IsValid()
	require.Equal(t, "The status must me enabled or disabled", err.Error())
	require.False(t, isValid)

	product.Status = application.ENABLED
	isValid, err = product.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)

	product.Price = -10
	isValid, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())
	require.False(t, isValid)
}
