package repositories

import (
	"wimf-services/entities"

	"gorm.io/gorm"
)

type FoodRepository interface {
	Save(food entities.Food) interface{}
	FindById(id int, userId string) interface{}
	FindByQuery(query, userId string) []entities.Food
	FindByUserId(userId string) []entities.Food
	Delete(id int, userId string) bool
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

func (f *foodRepository) FindById(id int, userId string) interface{} {
	var food entities.Food
	f.db.First(&food, "id = ? AND user_id = ?", id, userId)
	return food
}

func (f *foodRepository) FindByUserId(userId string) []entities.Food {
	var foods []entities.Food
	f.db.
		Find(&foods, "user_id = ?", userId).
		Order("id DESC")
	return foods
}

func (f *foodRepository) FindByQuery(query, userId string) []entities.Food {
	var foods []entities.Food
	res := f.db.
		Find(&foods, "("+
			"UPPER(name) LIKE UPPER(?) "+
			"OR UPPER(details) LIKE UPPER(?)"+
			") AND user_id = ?", "%"+query+"%", "%"+query+"%", userId).
		Order("id DESC")

	if res.Error != nil {
		return nil
	}
	return foods
}

func (f *foodRepository) Delete(id int, userId string) bool {
	res := f.db.
		Delete(&entities.Food{ID: id, UserID: userId})
	return res.Error == nil
}
