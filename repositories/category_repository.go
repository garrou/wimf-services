package repositories

import (
	"gorm.io/gorm"
	"wimf-services/entities"
)

type CategoryRepository interface {
	Get() []entities.Category
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Get() []entities.Category {
	var categories []entities.Category
	res := c.db.Find(&categories)

	if res.Error != nil {
		return nil
	}
	return categories
}
