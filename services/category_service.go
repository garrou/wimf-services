package services

import (
	"wimf-services/entities"
	"wimf-services/repositories"
)

type CategoryService interface {
	Get() []entities.Category
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (c *categoryService) Get() []entities.Category {
	return c.categoryRepository.Get()
}
