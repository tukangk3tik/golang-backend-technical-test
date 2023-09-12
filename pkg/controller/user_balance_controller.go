package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/dto/request/user"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/helper"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/service"
	"net/http"
)

type UserBalanceController interface {
	GetBalance(ctx *gin.Context)
	TopUpBalance(ctx *gin.Context)
	Transfer(ctx *gin.Context)
}

type balanceController struct {
	balanceService service.UserBalanceService
}

func NewUserBalanceController(balanceService service.UserBalanceService) UserBalanceController {
	return &balanceController{
		balanceService: balanceService,
	}
}

func (c *balanceController) GetBalance(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint)
	res, err := c.balanceService.GetUserBalance(userId)
	var response interface{} = helper.BuildSuccessResponse(res)

	if err != nil {
		response = helper.BuildFailResponse("Fail get balance", helper.EmptyObject{})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *balanceController) TopUpBalance(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint)

	var topUpDto user.TopUpBalanceDto
	errDto := ctx.ShouldBind(&topUpDto)
	if errDto != nil {
		response := helper.BuildFailResponse("Validation fail", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	header := helper.HeaderGetter(ctx)
	res, err := c.balanceService.TopUpBalance(userId, uint(topUpDto.Amount), header)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, helper.BuildSuccessResponse(res))
}

func (c *balanceController) Transfer(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint)

	var transferDto user.TransferDto
	errDto := ctx.ShouldBind(&transferDto)
	if errDto != nil {
		response := helper.BuildFailResponse("Validation fail", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	header := helper.HeaderGetter(ctx)
	res, err := c.balanceService.Transfer(userId, transferDto.To, uint(transferDto.Amount), header)
	if err != nil {
		ctx.JSON(http.StatusOK, helper.BuildFailResponse(err.Error(), helper.EmptyObject{}))
		return
	}

	ctx.JSON(http.StatusOK, helper.BuildSuccessResponse(res))
}
