package main

import (
	"net/http"

	"falabella.cl/prueba/api"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/beers", api.Beers)
	engine.POST("/beers", api.AddBeer)
	engine.GET("/beers/:id", api.BeerById)
	engine.GET("/beers/:id/boxprice", api.BoxPrice)

	http.ListenAndServe("0.0.0.0:8080", (engine))

}
