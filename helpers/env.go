package helpers

import "os"

var (
	Port     = os.Getenv("PORT")
	Username = os.Getenv("POSTGRES_USERNAME")
	Password = os.Getenv("POSTGRES_PASSWORD")
	DbName   = os.Getenv("POSTGRES_DB_NAME")
)
