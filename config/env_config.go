package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	projectPath := "/home/asad/go/src/github.com/AsaHero/simple-microservice/"
	err := godotenv.Load(projectPath + "/.env")
	if err != nil {
		panic(err)
	}
}
