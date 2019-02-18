package helpers

import "os"

var (
	API      = os.Getenv("COIN_API_KEY")
	Port     = os.Getenv("PORT")
	BASE_URL = os.Getenv("BASE_URL")
)
