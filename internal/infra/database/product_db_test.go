package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/dyhalmeida/go-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProductDb_ShoulInsertAProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)

	productDb := NewProductDb(db)
	err = productDb.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.GetID())
	assert.Equal(t, "Product 1", product.GetName())
}

func TestProductDb_ShouldFindAProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)

	productDb := NewProductDb(db)
	err = productDb.Create(product)
	assert.Nil(t, err)

	productFromDb, err := productDb.FindByID(product.GetID().String())
	assert.Nil(t, err)
	assert.Equal(t, product.GetID(), productFromDb.GetID())
	assert.Equal(t, product.GetName(), productFromDb.GetName())
	assert.Equal(t, product.GetPrice(), productFromDb.GetPrice())
}

func TestProductDb_ShouldFindAllProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	// Delete the products
	db.Exec("DELETE FROM products")

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}
	productDb := NewProductDb(db)
	products, err := productDb.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].GetName())
	assert.Equal(t, "Product 10", products[9].GetName())

	products, err = productDb.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].GetName())
	assert.Equal(t, "Product 20", products[9].GetName())

	products, err = productDb.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].GetName())
	assert.Equal(t, "Product 23", products[2].GetName())

}

func TestProductDb_ShouldDeleteAProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	db.Exec("DELETE FROM products")
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)

	productDb := NewProductDb(db)
	err = productDb.Create(product)
	assert.Nil(t, err)

	err = productDb.Delete(product.GetID().String())
	assert.Nil(t, err)

	_, err = productDb.FindByID(product.GetID().String())
	assert.NotNil(t, err)
}

func TestProductDb_ShouldBeUpdateAProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	db.Exec("DELETE FROM products")
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)

	productDb := NewProductDb(db)
	err = productDb.Create(product)
	assert.Nil(t, err)

	productFromDb, err := productDb.FindByID(product.GetID().String())
	assert.Nil(t, err)
	assert.Equal(t, product.GetID(), productFromDb.GetID())
	assert.Equal(t, "Product 1", productFromDb.GetName())
	assert.Equal(t, 10.5, productFromDb.GetPrice())

	product.Name = "Product 2"
	product.Price = 20.5
	err = productDb.Update(product)
	assert.Nil(t, err)

	productFromDb, err = productDb.FindByID(product.GetID().String())
	assert.Nil(t, err)
	assert.Equal(t, product.GetID(), productFromDb.GetID())
	assert.Equal(t, "Product 2", productFromDb.GetName())
	assert.Equal(t, 20.5, productFromDb.GetPrice())
}
