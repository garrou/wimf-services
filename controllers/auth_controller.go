package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wimf-services/dto"
	"wimf-services/entities"
	"wimf-services/helpers"
	"wimf-services/services"
)

type AuthController interface {
	Routes(e *gin.Engine)
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtHelper   helpers.JwtHelper
}

func NewAuthController(userService services.AuthService, jwtHelper helpers.JwtHelper) AuthController {
	return &authController{
		authService: userService,
		jwtHelper:   jwtHelper,
	}
}

func (a *authController) Routes(e *gin.Engine) {
	routes := e.Group("/api")
	{
		routes.POST("/register", a.Register)
		routes.POST("/login", a.Login)
	}
}

// Register creates user
func (a *authController) Register(ctx *gin.Context) {
	var userDto dto.UserCreateDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	userDto.TrimSpace()

	if !userDto.IsValid() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	if a.authService.IsDuplicateUsername(userDto.Username) {
		ctx.AbortWithStatusJSON(http.StatusConflict, helpers.NewResponse("Un nom d'utilisateur est déjà associé à ce compte", nil))
	} else {
		a.authService.Register(userDto)
		ctx.JSON(http.StatusCreated, helpers.NewResponse("Compte créé", nil))
	}
}

// Login authenticate user
func (a *authController) Login(ctx *gin.Context) {
	var userDto dto.UserLoginDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse("Données erronées", nil))
		return
	}
	res := a.authService.Login(userDto.Username, userDto.Password)

	if user, ok := res.(entities.User); ok {
		token := a.jwtHelper.GenerateToken(user.ID)
		ctx.JSON(http.StatusOK, helpers.NewResponse("", token))
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.NewResponse("Utilisateur ou mot de passe incorrect(s)", nil))
	}
}
