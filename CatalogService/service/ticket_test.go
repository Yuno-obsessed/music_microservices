package service_test

import (
	"context"
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/go-playground/assert/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"projects/music_microservices/StorageService/domain/dto"
	"projects/music_microservices/StorageService/domain/repository"
	"projects/music_microservices/StorageService/service"
)

func TestTicketService_GetEntity(t *testing.T) {
	db := setupTestDatabase()
	ticketService := service.TicketService{
		Repo: &repository.TicketRepository{
			Pool: db,
		},
	}
	query, args, _ := sq.Select("*").
		From("event_ticket").
		Where(sq.Eq{"event_id": 1}).ToSql()
	var expected dto.TicketOut
	_ = db.QueryRow(context.Background(), query, args).Scan(&expected)
	got, err := ticketService.GetEntity(1)
	if err != nil {
		t.Errorf("%v: \n%v", lerrors.ErrExecQuery, err)
	}
	assert.Equal(t, expected, got)
}

func setupTestDatabase() *pgxpool.Pool {
	// Set up test database
	db, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://user:password@localhost/test_db?sslmode=disable"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create tables and seed data
	_, err = db.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS event_ticket
	(
    	event_id INT PRIMARY KEY,
    	default_quantity INT,
    	vip_quantity INT,
    	scene_quantity INT,
    	default_cost INT NOT NULL,
    	vip_cost INT NOT NULL,
    	scene_cost INT NOT NULL,
    	FOREIGN KEY event_id
        	REFERENCES events (event_id)
        	ON DELETE CASCADE
	);

        INSERT INTO event_ticket (event_id, default_quantity, vip_quantity,
                                  scene_quantity, default_cost, vip_cost,
                                  scene_cost) VALUES
            (1, 200, 100, 140, 40, 60, 90),
            (2, 300, 160, 100, 30, 50, 80);
    `)
	if err != nil {
		panic(err)
	}

	return db
}
