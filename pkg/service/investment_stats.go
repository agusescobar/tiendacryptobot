package service

import (
	"encoding/json"
	d "github.com/agusfesc/tiendacryptobot/pkg/domain"
	"github.com/agusfesc/tiendacryptobot/pkg/utils"
	"log"
	"net/http"
)

func FindInvestmentStats() (*d.InvestmentStats, error) {
	resp, err := http.Get(utils.Enviroment.TiendaCryptoApi)
	if err != nil {
		log.Println("Error getting investment stats: ", err)
	}
	defer resp.Body.Close()

	investmentStats := d.InvestmentStats{}
	err = json.NewDecoder(resp.Body).Decode(&investmentStats)

	if err != nil {
		log.Println(err)
	}

	return &investmentStats, nil
}
