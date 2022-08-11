package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"wimf-services/controllers"
	"wimf-services/database"
	"wimf-services/helpers"
	"wimf-services/repositories"
	"wimf-services/services"
)

var (
	db        = database.Open()
	jwtHelper = helpers.NewJwtHelper()

	userRepository = repositories.NewUserRepository(db)
	authService    = services.NewAuthService(userRepository)
	authController = controllers.NewAuthController(authService, jwtHelper)

	userService    = services.NewUserService(userRepository)
	userController = controllers.NewUserController(userService, jwtHelper)

	categoryRepository = repositories.NewCategoryRepository(db)
	categoryService    = services.NewCategoryService(categoryRepository)
	categoryController = controllers.NewCategoryController(categoryService, jwtHelper)

	foodRepository = repositories.NewFoodRepository(db)
	foodService    = services.NewFoodService(foodRepository)
	foodController = controllers.NewFoodController(foodService, jwtHelper)
)

func main() {
	defer database.Close(db)

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err.Error())
	}
	authController.Routes(router)
	userController.Routes(router)
	categoryController.Routes(router)
	foodController.Routes(router)

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}
