package helpers

import "os"

var (
	Port     = os.Getenv("PORT")
	API      = os.Getenv("COIN_API_KEY")
	BASE_URL = os.Getenv("BASE_URL")
)
