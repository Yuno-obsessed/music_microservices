package main

import (
	"database-service/config"
	"database-service/services/database/connection"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")
	conf := config.ConfigInit()
	db := connection.NewDatabase(conf)
	go db.MigrationsRun()
	go db.Heartbeat()
}
