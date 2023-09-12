package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/helper"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/service"
	"net/http"
	"strings"
)

func AuthJwt(jwtService service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildFailResponse("No token. Please put your token", helper.EmptyObject{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		strToken := strings.Split(authHeader, " ")
		validateRes, err := jwtService.ValidateToken(strToken[1])
		if err != nil {
			response := helper.BuildFailResponse("Unauthorized", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("user_id", validateRes)
		c.Set("token", strToken[1])
		c.Next()
	}
}
