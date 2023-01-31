package connection

import (
	"context"
	"database-service/config"
	"database-service/services/logger"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Database struct {
	Pool   *pgxpool.Pool
	Logger logger.Logger
	Dns    string
}

func NewDatabase(conf config.Config) Database {
	return Database{
		Pool:   DbInit(conf),
		Logger: logger.NewLogger(conf),
		Dns:    DbDns(conf),
	}
}

func (d *Database) MigrationsRun() error {
	m, err := migrate.New("file://../../services/database/migrations", d.Dns)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func (d *Database) Heartbeat() {
	ticker := time.NewTicker(time.Second * 5)
	i := 0
	for range ticker.C {
		err := d.Pool.Ping(context.Background())
		if err != nil {
			d.Logger.Log.Error(fmt.Sprintf("Database connection was lost, %v", err))
			break
		}
		if i%5 == 0 {
			d.Logger.Log.Info("Database connection is alive")
		}
		i++
	}
}
