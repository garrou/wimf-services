package services

import (
	"time"
	"wimf-services/dto"
	"wimf-services/entities"
	"wimf-services/repositories"
)

type FoodService interface {
	Create(food dto.FoodCreateDto) interface{}
}

type foodService struct {
	foodRepository repositories.FoodRepository
}

func NewFoodService(foodRepository repositories.FoodRepository) FoodService {
	return &foodService{foodRepository: foodRepository}
}

func (f *foodService) Create(food dto.FoodCreateDto) interface{} {
	toCreate := entities.Food{
		Name:       food.Name,
		CategoryID: food.CategoryId,
		FreezeAt:   food.FreezeAt,
		Quantity:   food.Quantity,
		AddedAt:    time.Now(),
		UserID:     food.UserId,
	}
	res := f.foodRepository.Save(toCreate)

	if food, ok := res.(entities.Food); ok {
		return dto.FoodDto{
			Id:         food.ID,
			Name:       food.Name,
			Quantity:   food.Quantity,
			AddedAt:    food.AddedAt,
			FreezeAt:   food.FreezeAt,
			CategoryId: food.CategoryID,
		}
	}
	return nil
}
