package catalog

import (
	"github.com/joho/godotenv"
	"projects/music_microservices/StorageService/infra/server"
)

func main() {
	godotenv.Load("../../.env")
	server.Init()
}
