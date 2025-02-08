package entity

import (
	"errors"
	"time"

	"github.com/dyhalmeida/go-apis/pkg/entity"
)

var (
	ErrIdRequired    = errors.New("id is required")
	ErrIdInvalid     = errors.New("id is invalid")
	ErrNameRequired  = errors.New("name is required")
	ErrPriceRequired = errors.New("price is required")
	ErrPriceInvalid  = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil

}

func (p *Product) GetID() entity.ID {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *Product) Validate() error {

	if p.ID.String() == "" {
		return ErrIdRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIdInvalid
	}

	if p.Name == "" {
		return ErrNameRequired
	}

	if p.Price == 0 {
		return ErrPriceRequired
	}

	if p.Price < 0 {
		return ErrPriceInvalid
	}

	return nil

}
