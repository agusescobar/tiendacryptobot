package domain

type Enviroment struct {
	TiendaCryptoApi  string `json:"TIENDA_CRYPTO_API"`
	AustApi          string `json:"AUST_API"`
	TelegramApi      string `json:"TELEGRAM_API"`
	TelegramBotToken string `json:"TELEGRAM_BOT_TOKEN"`
}
