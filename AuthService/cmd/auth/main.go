package main

import (
	"auth-service/infra/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")
	server.Init()
}
