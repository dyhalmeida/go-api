package database

import (
	"github.com/dyhalmeida/go-apis/internal/entity"
	"gorm.io/gorm"
)

type ProductDbInterface interface {
	Create(*entity.Product) error
	FindByID(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}

type ProductDb struct {
	DB *gorm.DB
}

func NewProductDb(db *gorm.DB) *ProductDb {
	return &ProductDb{
		DB: db,
	}
}

func (u *ProductDb) Create(product *entity.Product) error {
	return u.DB.Create(product).Error
}

func (p *ProductDb) Update(product *entity.Product) error {
	if _, err := p.FindByID(product.ID.String()); err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *ProductDb) Delete(id string) error {
	if _, err := p.FindByID(id); err != nil {
		return err
	}
	return p.DB.Where("id = ?", id).Delete(&entity.Product{}).Error
}

func (p *ProductDb) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *ProductDb) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var product []entity.Product
	var err error

	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page <= 0 && limit <= 0 {
		err = p.DB.Order("created_at " + sort).Find(&product).Error
	} else {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&product).Error
	}

	return product, err
}
