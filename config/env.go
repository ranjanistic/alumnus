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
	SESSIONKEY string ""
	SITE string ""
	AUTH0SEC string ""
	DEV bool
}

var Env = Environment{
	ERR: godotenv.Load(),
	APPNAME: os.Getenv("APPNAME"),
	ENV: os.Getenv("ENV"),
	DBURL: os.Getenv("DBURL"),
	DBNAME: os.Getenv("DBNAME"),
	PORT: os.Getenv("PORT"),
	SESSIONKEY: os.Getenv("SESSIONKEY"),
	SITE:os.Getenv("SITE"),
	AUTH0SEC:os.Getenv("AUTH0SEC"),
	DEV: os.Getenv("ENV") != "production",
}
