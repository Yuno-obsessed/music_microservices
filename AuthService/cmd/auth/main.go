package main

import (
	"github.com/Yuno-obsessed/music_microservices/AuthService/infra/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")
	server.Init()
}
