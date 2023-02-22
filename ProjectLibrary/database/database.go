package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/config"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool   *pgxpool.Pool
	Logger logger.CustomLogger
	Dns    string
}

func NewDatabase() Database {
	return Database{
		Pool:   DbInit(),
		Logger: logger.NewLogger(),
		Dns:    DbDns(),
	}
}

func DbDns() string {
	conf := config.DatabaseConfigInit()
	return fmt.Sprintf("%s://%s:%s@localhost:%s/%s&parseTime=True",
		conf.Driver, conf.User,
		conf.Password, conf.Port, conf.Database)
}

func DbInit() *pgxpool.Pool {
	dns := DbDns()
	pool, err := pgxpool.New(context.Background(), dns)
	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
		return nil
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("database is unaccessible, %v", err)
		return nil
	}

	return pool
}
