package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func BindRequest[T any, U any](ctx *gin.Context, formStruct *T, uriStruct *U) bool {

	// validate uri
	if errUri := ctx.ShouldBindUri(&uriStruct); errUri != nil {
		log.Println(errUri)
		res := BuildFailResponse("NOT FOUND", EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
		return true
	}

	_ = ctx.ShouldBind(&formStruct)

	// validate query or body
	validate := validator.New()
	if errParam := validate.Struct(formStruct); errParam != nil {
		allErr := errParam.(validator.ValidationErrors)

		errList := make(map[string][]string)
		for _, errs := range allErr {
			fieldSnackCase := ToSnakeCase(errs.Field())
			errList[fieldSnackCase] = append(errList[fieldSnackCase], fmt.Sprintf("The %s error -> %s.", fieldSnackCase, errs.Tag()))
		}

		res := BuildFailResponse("Validation fail", errList)
		ctx.JSON(http.StatusBadRequest, res)
		return true
	}

	return false
}
