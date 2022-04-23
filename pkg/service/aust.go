package service

import (
	"encoding/json"
	d "github.com/agusfesc/tiendacryptobot/pkg/domain"
	"github.com/agusfesc/tiendacryptobot/pkg/utils"
	"log"
	"net/http"
)

func FindAustData() (*d.AUST, error) {
	resp, err := http.Get(utils.Enviroment.AustApi)
	if err != nil {
		log.Println("Error getting aust data: ", err)
	}
	defer resp.Body.Close()

	austData := d.AUST{}
	err = json.NewDecoder(resp.Body).Decode(&austData)

	if err != nil {
		log.Println(err)
	}

	return &austData, nil
}
