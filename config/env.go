package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Environment struct{
	ERR error
	APPNAME string ""
	ENV string ""
	DBURL string ""
	DBNAME string ""
	PORT string ""
}

var Env = Environment{
	ERR: godotenv.Load(),
	APPNAME: os.Getenv("APPNAME"),
	ENV: os.Getenv("ENV"),
	DBURL: os.Getenv("DBURL"),
	DBNAME: os.Getenv("DBNAME"),
	PORT: os.Getenv("PORT"),
}
