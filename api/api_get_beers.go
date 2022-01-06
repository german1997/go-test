package api

import (
	"net/http"

	"falabella.cl/prueba/data"
	"github.com/gin-gonic/gin"
)

func Beers(ctx *gin.Context) {
	r, e := getBeers(ctx)
	if e != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, &r)
}

func getBeers(ctx *gin.Context) (interface{}, error) {

	cervezas := data.Cervezas

	return cervezas, nil
}
