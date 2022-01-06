package model

type Beer struct {
	Id       int     `json:"id" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Brewery  string  `json:"brewery" binding:"required"`
	Country  string  `json:"country" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type BeerBox struct {
	PriceTotal float64 `json:"priceTotal"`
}

type Message struct {
	Message string `json:"message"`
}

type Rates struct {
	Success   bool        `json:"success"`
	Terms     string      `json:"terms"`
	Privacy   string      `json:"privacy"`
	Timestamp int         `json:"timestamp"`
	Source    string      `json:"source"`
	Quotes    interface{} `json:"quotes"`
}

type ValueCurrency struct {
	Data float64 `json:"data"`
}
