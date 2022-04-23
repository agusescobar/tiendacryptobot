package service

import (
	"fmt"
	u "github.com/agusfesc/tiendacryptobot/pkg/utils"
	"log"
	"net/http"
)

func SendMessage(channelId string, message string) error {
	var env = u.Enviroment
	path := fmt.Sprintf("%s/%s/sendMessage?chat_id=%s&text=%s", env.TelegramApi, env.TelegramBotToken, channelId, message)
	resp, err := http.Get(path)
	if err != nil {
		log.Println("Error sending message: ", err)
		return err
	}
	defer resp.Body.Close()
	return nil
}
