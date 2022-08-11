package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wimf-services/entities"
	"wimf-services/helpers"
	"wimf-services/middlewares"
	"wimf-services/services"
)

type UserController interface {
	Routes(e *gin.Engine)
	Get(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtHelper   helpers.JwtHelper
}

func NewUserController(userService services.UserService, jwtHelper helpers.JwtHelper) UserController {
	return &userController{userService: userService, jwtHelper: jwtHelper}
}

func (u *userController) Routes(e *gin.Engine) {
	routes := e.Group("/api/user", middlewares.AuthorizeJwt(u.jwtHelper))
	{
		routes.GET("/", u.Get)
	}
}

// Get gets the authenticated user
func (u *userController) Get(ctx *gin.Context) {
	userId := u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.Get(userId)

	if _, ok := res.(entities.User); ok {
		ctx.JSON(http.StatusOK, nil)
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de récupérer votre profil", nil))
	}
}
