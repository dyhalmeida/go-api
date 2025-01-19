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

type product struct {
	id        entity.ID
	name      string
	price     float64
	createdAt time.Time
}

func NewProduct(name string, price float64) (*product, error) {
	product := &product{
		id:        entity.NewID(),
		name:      name,
		price:     price,
		createdAt: time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil

}

func (p *product) GetID() entity.ID {
	return p.id
}

func (p *product) GetName() string {
	return p.name
}

func (p *product) GetPrice() float64 {
	return p.price
}

func (p *product) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *product) Validate() error {

	if p.id.String() == "" {
		return ErrIdRequired
	}

	if _, err := entity.ParseID(p.id.String()); err != nil {
		return ErrIdInvalid
	}

	if p.name == "" {
		return ErrNameRequired
	}

	if p.price == 0 {
		return ErrPriceRequired
	}

	if p.price < 0 {
		return ErrPriceInvalid
	}

	return nil

}
