package connection

import (
	"context"
	"database-service/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func DbDns(conf config.Config) string {
	return fmt.Sprintf("%s://%s:%s@localhost:%s/%s",
		conf.Database.Driver, conf.Database.User,
		conf.Database.Password, conf.Database.Port, conf.Database.Database)
}

func DbInit(conf config.Config) *pgxpool.Pool {
	dns := DbDns(conf)
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
