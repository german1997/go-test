package api

import (
	"strconv"

	"falabella.cl/prueba/data"
	"falabella.cl/prueba/model"
	"github.com/gin-gonic/gin"
)

func BeerById(ctx *gin.Context) {
	id := ctx.Param("id")
	r, codigo, e := getBeerById(ctx, id)
	if e != nil {
		ctx.AbortWithStatusJSON(codigo, r)
		return
	}
	ctx.JSON(codigo, &r)
}

func getBeerById(ctx *gin.Context, id string) (interface{}, int, error) {

	idBeer, _ := strconv.Atoi(id)

	var cerveza model.Beer
	for i := 0; i < len(data.Cervezas); i++ {
		if data.Cervezas[i].Id == idBeer {
			cerveza = data.Cervezas[i]
			i = len(data.Cervezas) + 1

			return cerveza, 200, nil
		}
	}

	var response model.Message
	response.Message = "El Id de la cerveza no existe"

	return response, 404, data.ErrorBeer
}
