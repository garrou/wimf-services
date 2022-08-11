package middlewares

import (
	"net/http"
	"wimf-services/helpers"

	"github.com/gin-gonic/gin"
)

func AuthorizeJwt(jwtHelper helpers.JwtHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := helpers.NewResponse("Aucun token", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtHelper.ValidateToken(authHeader)

		if !token.Valid || err != nil {
			response := helpers.NewResponse("Token invalide", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
