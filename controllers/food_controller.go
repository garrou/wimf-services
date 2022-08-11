package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wimf-services/dto"
	"wimf-services/helpers"
	"wimf-services/services"
)

type FoodController interface {
	Routes(e *gin.Engine)
	Create(ctx *gin.Context)
}

type foodController struct {
	foodService services.FoodService
	jwtHelper   helpers.JwtHelper
}

func NewFoodController(foodService services.FoodService, jwtHelper helpers.JwtHelper) FoodController {
	return &foodController{foodService: foodService, jwtHelper: jwtHelper}
}

func (f *foodController) Routes(e *gin.Engine) {
	routes := e.Group("/api/foods")
	{
		routes.POST("/", f.Create)
	}
}

func (f *foodController) Create(ctx *gin.Context) {
	var foodDto dto.FoodCreateDto

	if errDto := ctx.ShouldBind(&foodDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := f.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	foodDto.UserId = userId
	res := f.foodService.Create(foodDto)

	if food, ok := res.(dto.FoodDto); ok {
		response := helpers.NewResponse(fmt.Sprintf("%s ajouté(e) en quantité %d", food.Name, food.Quantity), food)
		ctx.JSON(http.StatusCreated, response)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Erreur durant l'ajout", nil))
	}
}
