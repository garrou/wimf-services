package repositories

import (
	"gorm.io/gorm"
	"wimf-services/entities"
)

type FoodRepository interface {
	Save(food entities.Food) interface{}
	Find() []entities.Food
}

type foodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &foodRepository{
		db: db,
	}
}

func (f *foodRepository) Save(food entities.Food) interface{} {
	res := f.db.Save(&food)

	if res.Error != nil {
		return nil
	}
	return food
}

func (f *foodRepository) Find() []entities.Food {
	var foods []entities.Food
	res := f.db.Find(foods)

	if res.Error != nil {
		return nil
	}
	return foods
}
