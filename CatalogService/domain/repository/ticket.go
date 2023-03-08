package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/multierr"
)

type TicketRepository struct {
	*pgxpool.Pool
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{
		database.DbInit(),
	}
}

func (t *TicketRepository) IsInCatalog(ttype string) (bool, error) {
	var amount int
	query, args, err := sq.Select("quantity").
		From("event_ticket").Where(sq.Eq{"type": ttype}).ToSql()
	if err != nil {
		return false, multierr.Append(lerrors.ErrInQuery, err)
	}
	row := t.Pool.QueryRow(context.Background(), query, args)
	err = row.Scan(&amount)
	if err != nil {
		return false, multierr.Append(lerrors.ErrNoRecord, err)
	}
	if amount <= 0 {
		return false, fmt.Errorf("no tickets are available")
	}
	return true, nil
}

func (t *TicketRepository) GetQuantityOfType(ttype string) (int, error) {
	var amount int
	query, args, err := sq.Select("quantity").
		From("event_ticket").Where(sq.Eq{"type": ttype}).ToSql()
	if err != nil {
		return 0, multierr.Append(lerrors.ErrInQuery, err)
	}
	row := t.Pool.QueryRow(context.Background(), query, args)
	err = row.Scan(&amount)
	if err != nil {
		return 0, multierr.Append(lerrors.ErrNoRecord, err)
	}
	return amount, nil
}

func (t *TicketRepository) Subtruct(ttype string, number int) error {
	prevAmount, err := t.GetQuantityOfType(ttype)
	if err != nil {
		return err
	}
	query, args, err := sq.Update("event_ticket").
		Set("amount", prevAmount-number).
		Where(sq.Eq{"type": ttype}).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}
	_, err = t.Pool.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}

	return nil
}
