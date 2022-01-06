package model

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorApiResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func GlobalError(ctx *gin.Context, errors error, model interface{}) {
	errs := errors.(validator.ValidationErrors)
	var errorResponse []ErrorApiResponse
	var fieldName string
	hasError := false
	for _, err := range errs {
		hasError = true
		t := reflect.TypeOf(model)
		field, find := t.FieldByName(err.Field())
		if !find {
			fieldName = err.Field()
		} else {
			fieldName = getTag(ctx.Request.Method, field)
		}
		errorResponse = append(errorResponse,
			ErrorApiResponse{
				Type:    err.Tag(),
				Message: SetMessageByClassification(err.Tag(), fieldName),
			})

	}
	if hasError {
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}
}

func SetMessageByClassification(classification, field string) string {
	switch classification {
	case "required":
		return fmt.Sprintf("Missing required field %s", field)
	case "min":
		return fmt.Sprintf("Invalid size in %s", field)
	case "max":
		return fmt.Sprintf("Invalid size in %s", field)
	default:
		return fmt.Sprintf("Bad value in field %s", field)
	}
}

func getTag(method string, field reflect.StructField) string {
	switch method {
	case http.MethodPost:
		return field.Tag.Get("json")
	case http.MethodDelete:
		return field.Tag.Get("json")
	case http.MethodGet:
		return field.Tag.Get("form")
	}
	return field.Name
}
