package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/dto"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/EventService/service/event/interfaces"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/multierr"
)

type EventRepository struct {
	*pgxpool.Pool
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		database.DbInit(),
	}
}

var _ interfaces.EventInterface = EventRepository{}

func (e *EventRepository) GetOne(id int) (dto.Event, error) {
	var event dto.Event
	query, args, err := sq.Select("band_name", "event_city").
		From("events").Where(sq.Eq{"event_id": id}).ToSql()
	if err != nil {
		return dto.Event{}, multierr.Append(lerrors.ErrInQuery, err)
	}
	row := e.Pool.QueryRow(context.Background(), query, args)
	err = row.Scan(&event)
	if err != nil {
		return dto.Event{}, multierr.Append(lerrors.ErrNoRecord, err)
	}

	return event, nil
}

func (e *EventRepository) GetAllOfBand(band string) ([]entity.Event, error) {
	var events []entity.Event
	query, args, err := sq.Select("*").
		From("events").Where(sq.Eq{"event_band": band}).ToSql()
	if err != nil {
		return nil, multierr.Append(lerrors.ErrInQuery, err)
	}
	rows, err := e.Pool.Query(context.Background(), query, args)
	if err != nil {
		return nil, multierr.Append(lerrors.ErrNoRecord, err)
	}
	defer rows.Close()
	index := 0
	for rows.Next() {
		curr := events[index]
		err = rows.Scan(&curr.EventId, &curr.BandName,
			&curr.EventCity, &curr.Date)
		if err != nil {
			return nil, multierr.Append(lerrors.ErrNoRecord, err)
		}
	}

	return events, nil
}

func (e *EventRepository) GetAllOfCity(city string) ([]entity.Event, error) {
	var events []entity.Event
	query, args, err := sq.Select("*").
		From("events").Where(sq.Eq{"event_city": city}).ToSql()
	if err != nil {
		return nil, multierr.Append(lerrors.ErrInQuery, err)
	}
	rows, err := e.Pool.Query(context.Background(), query, args)
	if err != nil {
		return nil, multierr.Append(lerrors.ErrNoRecord, err)
	}
	defer rows.Close()
	index := 0
	for rows.Next() {
		curr := events[index]
		err = rows.Scan(&curr.EventId, &curr.BandName,
			&curr.EventCity, &curr.Date)
		if err != nil {
			return nil, multierr.Append(lerrors.ErrNoRecord, err)
		}
	}

	return events, nil
}

// add limit and offset
func (e *EventRepository) GetAll() ([]entity.Event, error) {
	var events []entity.Event
	query, args, err := sq.Select("*").
		From("events").ToSql()
	if err != nil {
		return nil, multierr.Append(lerrors.ErrInQuery, err)
	}
	rows, err := e.Pool.Query(context.Background(), query, args)
	if err != nil {
		return nil, multierr.Append(lerrors.ErrNoRecord, err)
	}
	defer rows.Close()
	index := 0
	for rows.Next() {
		curr := events[index]
		err = rows.Scan(&curr.EventId, &curr.BandName,
			&curr.EventCity, &curr.Date)
		if err != nil {
			return nil, multierr.Append(lerrors.ErrNoRecord, err)
		}
	}

	return events, nil
}

func (e *EventRepository) Create(event dto.Event) (int, error) {
	var res int
	query, args, err := sq.Insert("events").
		Columns("event_band", "event_city", "event_date").
		Values(event.BandName, event.EventCity, event.Date).
		Suffix("RETURNING event_id").ToSql()
	if err != nil {
		return 0, multierr.Append(lerrors.ErrInQuery, err)
	}
	err = e.Pool.QueryRow(context.Background(), query, args).Scan(&res)
	if err != nil {
		return 0, multierr.Append(lerrors.ErrExecQuery, err)
	}

	return res, nil
}

func (e *EventRepository) Update(id int, event dto.Event) error {
	query, args, err := sq.Update("events").
		Set("event_band", event.BandName).
		Set("event_city", event.EventCity).
		Set("event_date", event.Date).
		Where(sq.Eq{"event_id": id}).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}
	_, err = e.Pool.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}

	return nil
}

func (e *EventRepository) Delete(id int) error {
	query, args, err := sq.Delete("events").
		Where(sq.Eq{"event_id": id}).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}
	_, err = e.Pool.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}

	return nil
}
