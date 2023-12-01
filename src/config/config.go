package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StrConn   = ""
	Port      = 0
	SecretKey []byte
)

func LoadEnvs() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("PORT"))
	if erro != nil {
		Port = 9000
	}

	StrConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DB"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
