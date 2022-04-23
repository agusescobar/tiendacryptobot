package main

import (
	"fmt"
	"github.com/agusfesc/tiendacryptobot/pkg/domain"
	"github.com/agusfesc/tiendacryptobot/pkg/service"
	"github.com/agusfesc/tiendacryptobot/pkg/utils"
	"log"
	"time"
)

const SPACE = "%0A"

func main() {
	utils.LoadEnviroment()

	for range time.Tick(time.Minute * 30) {
		isOk := processData()
		if !isOk {
			log.Println("Error getting data")
		} else {
			log.Println("Data processed successfully")
		}
	}
}

func processData() bool {
	austData, err := service.FindAustData()
	if err != nil {
		log.Println("Error executing FindAustData service")
		return false
	}
	investmentStats, err := service.FindInvestmentStats()
	if err != nil {
		log.Println("Error executing FindInvestmentStats service")
		return false
	}

	message := getMessage(investmentStats, austData)

	err = service.SendMessage("-1001620082911", message)
	if err != nil {
		log.Println("Error executing SendMessage service")
		return false
	}

	return true
}

func getMessage(investmentStats *domain.InvestmentStats, austData *domain.AUST) string {
	austPrice := austData.Prices.Coin.Price
	investmentPercentage := investmentStats.Real * 100
	message := fmt.Sprintf("El precio del aUST es de: $%.4f USD %sEl porcentaje de inversion actual es de: %.4f%% ", austPrice, SPACE, investmentPercentage)
	return message
}
