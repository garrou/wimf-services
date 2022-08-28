package repositories

import (
	"wimf-services/entities"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Find() []entities.Category
	FindFoodsByCategory(id int, userId string) []entities.Food
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Find() []entities.Category {
	var categories []entities.Category
	res := c.db.
		Find(&categories).
		Order("id")

	if res.Error != nil {
		return nil
	}
	return categories
}

func (c *categoryRepository) FindFoodsByCategory(id int, userId string) []entities.Food {
	var foods []entities.Food
	res := c.db.
		Find(&foods, "category_id = ? AND user_id = ?", id, userId).
		Order("id DESC")

	if res.Error != nil {
		return nil
	}
	return foods
}
