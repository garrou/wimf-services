package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wimf-services/dto"
	"wimf-services/helpers"
	"wimf-services/services"
)

type CategoryController interface {
	Routes(e *gin.Engine)
	Get(ctx *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
	jwtHelper       helpers.JwtHelper
}

func NewCategoryController(categoryService services.CategoryService, jwtHelper helpers.JwtHelper) CategoryController {
	return &categoryController{categoryService: categoryService, jwtHelper: jwtHelper}
}

func (c *categoryController) Routes(e *gin.Engine) {
	routes := e.Group("/api/categories")
	{
		routes.GET("/", c.Get)
	}
}

func (c *categoryController) Get(ctx *gin.Context) {
	var categories []dto.CategoryDto
	res := c.categoryService.Get()

	for _, c := range res {
		categories = append(categories, dto.CategoryDto{
			Id:    c.ID,
			Name:  c.Name,
			Image: c.Image,
		})
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse("", categories))
}
