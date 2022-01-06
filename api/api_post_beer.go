package api

import (
	"falabella.cl/prueba/data"
	"falabella.cl/prueba/model"

	"github.com/gin-gonic/gin"
)

func AddBeer(ctx *gin.Context) {
	var request model.Beer
	if err := ctx.ShouldBind(&request); err != nil {
		model.GlobalError(ctx, err, model.Beer{})
		return
	}

	r, codigo, e := getAddBeer(ctx, request)
	if e != nil {
		ctx.AbortWithStatusJSON(codigo, r)
		return
	}
	ctx.JSON(codigo, &r)
}

func getAddBeer(ctx *gin.Context, request model.Beer) (interface{}, int, error) {

	var existe bool
	var response model.Message

	for i := 0; i < len(data.Cervezas); i++ {
		if data.Cervezas[i].Id == request.Id {
			existe = true
			i = len(data.Cervezas) + 1
		}
	}

	if existe {
		response.Message = "El ID de la cerveza ya existe"
		return response, 409, data.ErrorBeer
	} else {
		data.Cervezas = append(data.Cervezas, request)
		response.Message = "Cerveza creada"
	}

	return response, 201, nil
}
