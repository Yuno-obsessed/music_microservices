package database

import (
	"auth-service/infra/config"
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DbInit() (*pgxpool.Pool, error) {
	conf := config.DatabaseConfigInit()
	dns := fmt.Sprintf("%s://%s:%s@localhost:%s/%s",
		conf.Driver, conf.User, conf.Password, conf.Port, conf.Database)

	pool, err := pgxpool.New(context.Background(), dns)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database, %v", err)
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("database is unaccessible, %v", err)
	}

	m, err := migrate.New("file://migrations", dns)
	if err != nil {
		return nil, fmt.Errorf("error linking migrations to db, %v", err)
	}

	if err := m.Up(); err != nil {
		return nil, fmt.Errorf("error performing migrate up, %v", err)
	}

	return pool, nil
}
