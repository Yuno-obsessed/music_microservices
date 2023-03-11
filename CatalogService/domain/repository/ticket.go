package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/multierr"
	"projects/music_microservices/StorageService/domain/dto"
	"projects/music_microservices/StorageService/service/catalog/interfaces"
)

type TicketRepository struct {
	*pgxpool.Pool
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{
		database.DbInit(),
	}
}

var _ interfaces.TicketInterface = &TicketRepository{}

func (t *TicketRepository) GetSumAndAverage(id int) (dto.TicketBrief, error) {
	var ticket dto.TicketBrief
	// map to struct
	query, args, err := sq.Select("COALESCE(default_quantity, 0) + COALESCE(vip_quantity, 0) + "+
		"COALESCE(scene_quantity, 0) AS total_amount", "AVG(default_cost, vip_cost, scene_cost)").
		From("event_ticket").Where(sq.Eq{"event_id": id}).ToSql()
	if err != nil {
		return dto.TicketBrief{}, multierr.Append(lerrors.ErrInQuery, err)
	}
	err = t.Pool.QueryRow(context.Background(), query, args).Scan(&ticket)
	if err != nil {
		return dto.TicketBrief{}, multierr.Append(lerrors.ErrNoRecord, err)
	}
	if ticket.TicketsCount == 0 {
		return dto.TicketBrief{}, fmt.Errorf("no tickets are available")
	}

	return ticket, nil
}

func (t *TicketRepository) GetEntity(id int) (dto.TicketOut, error) {
	var event dto.TicketOut
	query, args, err := sq.Select("default_quantity", "vip_quantity",
		"scene_quantity", "default_cost", "vip_cost", "scene_cost").
		From("event_ticket").Where(sq.Eq{"event_id": id}).ToSql()
	if err != nil {
		return dto.TicketOut{}, multierr.Append(lerrors.ErrInQuery, err)
	}
	err = t.Pool.QueryRow(context.Background(), query, args).Scan(&event)
	if err != nil {
		return dto.TicketOut{}, multierr.Append(lerrors.ErrNoRecord, err)
	}

	return event, nil
}

func (t *TicketRepository) Subtruct(id int, ttype consts.TicketType, number int) error {
	var ticket string
	switch ttype {
	case consts.DEFAULT:
		ticket = "default_quantity"
	case consts.VIP:
		ticket = "vip_quantity"
	case consts.SCENE:
		ticket = "scene_quantity"
	}
	query, args, err := sq.Select(ticket).From("event_ticket").
		Where(sq.Eq{"event_id": id}).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}
	var quantity int
	err = t.Pool.QueryRow(context.Background(), query, args).Scan(&quantity)
	if err != nil {
		return multierr.Append(lerrors.ErrNoRecord, err)
	}
	query, args, err = sq.Update("event_ticket").
		Set(ticket, quantity-number).
		Where(sq.Eq{"event_id": id}).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}
	_, err = t.Pool.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}

	return nil
}
