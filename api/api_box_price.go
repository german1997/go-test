package api

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"falabella.cl/prueba/data"
	"falabella.cl/prueba/model"
	"github.com/gin-gonic/gin"
)

func BoxPrice(ctx *gin.Context) {

	id := ctx.Param("id")
	currency := ctx.Query("currency")
	quantity := ctx.Query("quantity")

	r, codigo, e := getBoxPrice(ctx, id, currency, quantity)
	if e != nil {
		ctx.AbortWithStatusJSON(codigo, r)
		return
	}
	ctx.JSON(codigo, &r)
}

func getBoxPrice(ctx *gin.Context, id string, currency string, quantity string) (interface{}, int, error) {

	idBeer, _ := strconv.Atoi(id)
	quantityBeer, _ := strconv.ParseFloat(quantity, 64)

	var beerBox model.BeerBox
	var total float64
	var existe bool

	for i := 0; i < len(data.Cervezas); i++ {
		if data.Cervezas[i].Id == idBeer {
			existe = true
			total = data.Cervezas[i].Price * quantityBeer
			i = len(data.Cervezas) + 1
		}
	}

	if existe {

		access_key := "e09ac34b400a407731367781f377564c"
		URL := "http://api.currencylayer.com/live?" + "access_key=" + access_key
		resp, err := http.Get(URL)

		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		var rates model.Rates
		json.Unmarshal(body, &rates)

		valueQuotes := rates.Quotes.(map[string]interface{})
		ValueCurrency := valueQuotes["USD"+currency].(float64)

		priceConverted := (total * ValueCurrency) / 1
		beerBox.PriceTotal = toFixed(priceConverted, 2)

		return beerBox, 200, nil
	}

	var response model.Message
	response.Message = "El Id de la cerveza no existe"

	return response, 404, nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
