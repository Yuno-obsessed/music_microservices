package main

import (
	"auth-service/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")
	server.Init()
}
