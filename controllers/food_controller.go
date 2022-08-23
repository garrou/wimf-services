package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wimf-services/dto"
	"wimf-services/helpers"
	"wimf-services/services"
)

type FoodController interface {
	Routes(e *gin.Engine)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Search(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
		routes.GET("/", f.Get)
		routes.GET("/search", f.Search)
		routes.PUT("/", f.Update)
		routes.DELETE("/:id", f.Delete)
	}
}

func (f *foodController) Create(ctx *gin.Context) {
	var foodDto dto.FoodCreateDto

	if errDto := ctx.ShouldBind(&foodDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	foodDto.UserId = f.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := f.foodService.Create(foodDto)

	if food, ok := res.(dto.FoodDto); ok {
		response := helpers.NewResponse(fmt.Sprintf("%s ajouté(e) en quantité %d", food.Name, food.Quantity), food)
		ctx.JSON(http.StatusCreated, response)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Erreur durant l'ajout", nil))
	}
}

func (f *foodController) Get(ctx *gin.Context) {
	userId := f.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	foods := f.foodService.GetByUserId(userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", foods))
}

func (f *foodController) Update(ctx *gin.Context) {
	var foodDto dto.FoodUpdateDto

	if errDto := ctx.ShouldBind(&foodDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	foodDto.UserId = f.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := f.foodService.Update(foodDto)

	if food, ok := res.(dto.FoodDto); ok {
		response := helpers.NewResponse(fmt.Sprintf("%s modifié(e)(s)", food.Name), food)
		ctx.JSON(http.StatusCreated, response)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Erreur durant la modification", nil))
	}
}

func (f *foodController) Search(ctx *gin.Context) {
	query := ctx.Query("q")
	userId := f.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	foods := f.foodService.Search(query, userId)
	ctx.JSON(http.StatusOK, helpers.NewResponse("", foods))
}

func (f *foodController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userId := f.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	deleted := f.foodService.Delete(id, userId)

	if deleted {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Aliment supprimé", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Erreur durant la suppression", nil))
	}
}
