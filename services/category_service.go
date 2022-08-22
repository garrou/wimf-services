package services

import (
	"wimf-services/dto"
	"wimf-services/repositories"
)

type CategoryService interface {
	Get() []dto.CategoryDto
	GetFoodsByCategory(id int, userId string) []dto.FoodDto
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (c *categoryService) Get() []dto.CategoryDto {
	var categories []dto.CategoryDto
	res := c.categoryRepository.Get()

	for _, c := range res {
		categories = append(categories, dto.CategoryDto{
			Id:    c.ID,
			Name:  c.Name,
			Image: c.Image,
		})
	}
	return categories
}

func (c *categoryService) GetFoodsByCategory(id int, userId string) []dto.FoodDto {
	var foods []dto.FoodDto
	res := c.categoryRepository.GetFoodsByCategory(id, userId)

	for _, f := range res {
		foods = append(foods, dto.FoodDto{
			Id:         f.ID,
			Name:       f.Name,
			Details:    f.Details,
			CategoryId: f.CategoryID,
			AddedAt:    f.AddedAt,
			Quantity:   f.Quantity,
			FreezeAt:   f.FreezeAt,
		})
	}
	return foods
}
