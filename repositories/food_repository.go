package repositories

import (
	"gorm.io/gorm"
	"wimf-services/entities"
)

type FoodRepository interface {
	Save(food entities.Food) interface{}
	Find() []entities.Food
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

func (f *foodRepository) Find() []entities.Food {
	var foods []entities.Food
	res := f.db.Find(&foods)

	if res.Error != nil {
		return nil
	}
	return foods
}

func (f *foodRepository) FindById(id int, userId string) interface{} {
	var food entities.Food
	f.db.
		Find(&food, "id = ? AND user_id = ?", id, userId).
		Order("added_at DESC")
	return food
}

func (f *foodRepository) FindByUserId(userId string) []entities.Food {
	var foods []entities.Food
	f.db.
		Find(&foods, "user_id = ?", userId).
		Order("added_at DESC")
	return foods
}

func (f *foodRepository) FindByQuery(query, userId string) []entities.Food {
	var foods []entities.Food
	res := f.db.
		Find(&foods, "("+
			"UPPER(name) LIKE UPPER(?) "+
			"OR UPPER(details) LIKE UPPER(?)"+
			") AND user_id = ?", "%"+query+"%", "%"+query+"%", userId).
		Order("added_at DESC")

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
