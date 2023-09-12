package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/dto/request/auth"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/entity"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/helper"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/service"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JwtService
}

func NewAuthController(authService service.AuthService, jwtService service.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto auth.LoginDto
	errDto := ctx.ShouldBind(&loginDto)
	if errDto != nil {
		response := helper.BuildFailResponse("Validation fail", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.Login(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(v.ID)
		response := helper.BuildSuccessResponse(generateToken)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildFailResponse("Invalid credential", helper.EmptyObject{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Logout(ctx *gin.Context) {
	strToken := ctx.MustGet("token").(string)
	err := c.jwtService.RemoveToken(strToken)
	if err != nil {
		res := helper.BuildFailResponse("Invalid token", helper.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusOK, helper.BuildSuccessResponse(helper.EmptyObject{}))
}
