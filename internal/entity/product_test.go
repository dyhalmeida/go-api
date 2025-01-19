package entity

import (
	"testing"

	"github.com/dyhalmeida/go-apis/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestProduct_ShouldBeInvalidName(t *testing.T) {
	product, err := NewProduct("", 10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameRequired, err)
	assert.Nil(t, product)
}

func TestProduct_ShouldBeInvalidPrice(t *testing.T) {
	product, err := NewProduct("Product", -10)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceInvalid, err)
	assert.Nil(t, product)
}

func TestProduct_ShouldBeRequiredPrice(t *testing.T) {
	product, err := NewProduct("Product", 0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceRequired, err)
	assert.Nil(t, product)
}

func TestProduct_ShouldBeValidProduct(t *testing.T) {
	product, err := NewProduct("Product", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.GetID())
	assert.Equal(t, "Product", product.GetName())
	assert.Equal(t, 10.0, product.GetPrice())
}

func TestProduct_ShoulCallValidateSuccessfully(t *testing.T) {
	product := &product{
		id:    entity.NewID(),
		name:  "Product",
		price: 10.0,
	}
	assert.Nil(t, product.Validate())
}
