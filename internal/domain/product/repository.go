package product

import "gorm.io/gorm"

type (
	ProductRepository interface {
		Create(ProductSchema) error
		Update(string, ProductSchema) error
		GetByID(string) (ProductSchema, error)
		Delete(string) error
	}

	productRepo struct {
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(product ProductSchema) error {
	if err := r.db.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) Update(id string, product ProductSchema) error {
	if err := r.db.Model(&ProductSchema{}).Where("id = ?", id).Updates(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) GetByID(id string) (ProductSchema, error) {
	var product ProductSchema
	if err := r.db.First(&product, id).Error; err != nil {
		return ProductSchema{}, err
	}

	return product, nil
}

func (r *productRepo) Delete(id string) error {
	if err := r.db.Delete(&ProductSchema{}, id).Error; err != nil {
		return err
	}

	return nil
}
