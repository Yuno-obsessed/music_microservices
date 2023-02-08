package event

import "github.com/jackc/pgx/v5/pgxpool"

type EventService struct {
	Repo interface{}
	Pool *pgxpool.Pool
}
