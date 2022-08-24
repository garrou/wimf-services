package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wimf-services/helpers"
	"wimf-services/middlewares"
	"wimf-services/services"
)

type CategoryController interface {
	Routes(e *gin.Engine)
	Get(ctx *gin.Context)
	GetFoodsByCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
	jwtHelper       helpers.JwtHelper
}

func NewCategoryController(categoryService services.CategoryService, jwtHelper helpers.JwtHelper) CategoryController {
	return &categoryController{categoryService: categoryService, jwtHelper: jwtHelper}
}

func (c *categoryController) Routes(e *gin.Engine) {
	routes := e.Group("/api/categories", middlewares.AuthorizeJwt(c.jwtHelper))
	{
		routes.GET("/", c.Get)
		routes.GET("/:id/foods", c.GetFoodsByCategory)
	}
}

func (c *categoryController) Get(ctx *gin.Context) {
	categories := c.categoryService.Get()
	ctx.JSON(http.StatusOK, helpers.NewResponse("", categories))
}

func (c *categoryController) GetFoodsByCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := c.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	foods := c.categoryService.GetFoodsByCategory(id, userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", foods))
}
