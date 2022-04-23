package domain

type AUST struct {
	Prices `json:"prices"`
}

type Prices struct {
	Coin `json:"aUST"`
}

type Coin struct {
	Price        float64 `json:"price"`
	PctChange1h  float64 `json:"pct_change_1h"`
	PctChange24h float64 `json:"pct_change_24h"`
	PctChange7d  float64 `json:"pct_change_7d"`
	MarketCap    float64 `json:"market_cap"`
}
