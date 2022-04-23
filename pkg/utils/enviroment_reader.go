package utils

import (
	"bytes"
	"encoding/json"
	"github.com/agusfesc/tiendacryptobot/pkg/domain"
	"github.com/joho/godotenv"
	"log"
)

var Enviroment *domain.Enviroment

func LoadEnviroment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	enviroment, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	transcode(enviroment, &Enviroment)
}

func transcode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
