package services

import (
	"wimf-services/dto"
	"wimf-services/entities"
	"wimf-services/repositories"
)

type FoodService interface {
	Create(food dto.FoodCreateDto) interface{}
	Update(food dto.FoodUpdateDto) interface{}
	Search(query, userId string) []dto.FoodDto
	GetByUserId(userId string) []dto.FoodDto
	Delete(id int, userId string) bool
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
		Details:    food.Details,
		UserID:     food.UserId,
	}
	res := f.foodRepository.Save(toCreate)

	if food, ok := res.(entities.Food); ok {
		return dto.FoodDto{
			Id:         food.ID,
			Name:       food.Name,
			Quantity:   food.Quantity,
			FreezeAt:   food.FreezeAt,
			Details:    food.Details,
			CategoryId: food.CategoryID,
		}
	}
	return nil
}

func (f *foodService) Update(foodDto dto.FoodUpdateDto) interface{} {
	res := f.foodRepository.FindById(foodDto.Id, foodDto.UserId)

	if food, ok := res.(entities.Food); ok {
		food.Name = foodDto.Name
		food.Quantity = foodDto.Quantity
		food.FreezeAt = foodDto.FreezeAt
		food.Details = foodDto.Details

		res = f.foodRepository.Save(food)

		if food, ok := res.(entities.Food); ok {
			return dto.FoodDto{
				Id:         food.ID,
				Name:       food.Name,
				Quantity:   food.Quantity,
				FreezeAt:   food.FreezeAt,
				Details:    food.Details,
				CategoryId: food.CategoryID,
			}
		}
	}
	return nil
}

func (f *foodService) Search(query, userId string) []dto.FoodDto {
	var foods []dto.FoodDto
	res := f.foodRepository.FindByQuery(query, userId)

	for _, f := range res {
		foods = append(foods, dto.FoodDto{
			Id:         f.ID,
			Name:       f.Name,
			Quantity:   f.Quantity,
			FreezeAt:   f.FreezeAt,
			Details:    f.Details,
			CategoryId: f.CategoryID,
		})
	}
	return foods
}

func (f *foodService) GetByUserId(userId string) []dto.FoodDto {
	var foods []dto.FoodDto
	res := f.foodRepository.FindByUserId(userId)

	for _, f := range res {
		foods = append(foods, dto.FoodDto{
			Id:         f.ID,
			Name:       f.Name,
			Quantity:   f.Quantity,
			FreezeAt:   f.FreezeAt,
			Details:    f.Details,
			CategoryId: f.CategoryID,
		})
	}
	return foods
}

func (f *foodService) Delete(id int, userId string) bool {
	return f.foodRepository.Delete(id, userId)
}
