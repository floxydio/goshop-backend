package repository

import (
	"goshop/database"
	models "goshop/internal/model"

	"gorm.io/gorm"
)

type ProductDB struct {
	PDB *gorm.DB
}

type ProductRepo interface {
	FindProduct(string, string) ([]models.Product, error)
	CreateProduct(form models.ProductForm) error
	UpdateProduct(form models.ProductForm, id uint) error
	DeleteProduct(id uint) error
}

func ProductRepoIntf(db *gorm.DB) ProductRepo {
	return &ProductDB{
		PDB: database.DbClient,
	}
}

func (p *ProductDB) FindProduct(name string, category string) ([]models.Product, error) {
	var err error
	var products []models.Product

	if name != "" {
		err = p.PDB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			return nil, err
		}
	}
	if category != "" {
		err = p.PDB.Where("category LIKE ?", "%"+category+"%").Find(&products).Error
		if err != nil {
			return nil, err
		}
	}
	if name == "" && category == "" {
		err = p.PDB.Find(&products).Error
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

func (p *ProductDB) CreateProduct(form models.ProductForm) error {
	err := p.PDB.Create(form).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDB) UpdateProduct(form models.ProductForm, id uint) error {
	err := p.PDB.Model(&models.Product{}).Where("product_id = ?", id).Updates(form).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDB) DeleteProduct(id uint) error {
	err := p.PDB.Where("product_id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
