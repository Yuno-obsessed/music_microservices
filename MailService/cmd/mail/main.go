package main

import (
	"github.com/joho/godotenv"
	"mail-service/infra/server"
)

func main() {
	godotenv.Load("../../.env")
	server.Init()
}
