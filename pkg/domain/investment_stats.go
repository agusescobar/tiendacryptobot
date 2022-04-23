package domain

type InvestmentStats struct {
	Theoric    float64 `json:"theoric"`
	Real       float64 `json:"real"`
	Fee        float64 `json:"fee"`
	Multiplier float64 `json:"multiplier"`
}
