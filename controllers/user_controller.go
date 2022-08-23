package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wimf-services/dto"
	"wimf-services/entities"
	"wimf-services/helpers"
	"wimf-services/middlewares"
	"wimf-services/services"
)

type UserController interface {
	Routes(e *gin.Engine)
	Get(ctx *gin.Context)
	UpdateUsername(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
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
		routes.PATCH("/username", u.UpdateUsername)
		routes.PATCH("/password", u.UpdatePassword)
	}
}

// Get gets the authenticated user
func (u *userController) Get(ctx *gin.Context) {
	userId := u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.Get(userId)

	if user, ok := res.(entities.User); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("", dto.UserDto{
			Username: user.Username,
		}))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de récupérer votre profil", nil))
	}
}

func (u *userController) UpdateUsername(ctx *gin.Context) {
	var usernameDto dto.UsernameDto

	if errDto := ctx.ShouldBind(&usernameDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	usernameDto.UserId = u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.UpdateUsername(usernameDto)

	if _, ok := res.(entities.User); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Identifiant modifié", nil))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Impossible de modifier votre identifiant", nil))
	}
}

func (u *userController) UpdatePassword(ctx *gin.Context) {
	var userDto dto.PasswordDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userDto.UserId = u.jwtHelper.ExtractUserId(ctx.GetHeader("Authorization"))
	res := u.userService.UpdatePassword(userDto)

	if _, ok := res.(entities.User); ok {
		ctx.JSON(http.StatusOK, helpers.NewResponse("Mot de passe modifié", nil))
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.NewResponse("Impossible de modifier votre mot de passe", nil))
	}
}
