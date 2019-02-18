package helpers

import "os"

var (
	API = os.Getenv("COIN_API_KEY")
)

func SetBASE() string {
	BASE_URL := os.Getenv("BASE_URL")
	if BASE_URL == "" {
		BASE_URL = "https://pro-api.coinmarketcap.com/v1"
	}
	return BASE_URL
}

func SetPort() string {
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	return Port
}
