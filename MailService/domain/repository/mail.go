package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/MailService/domain/dto"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/multierr"
)

type MailRepository struct {
	*pgxpool.Pool
}

func NewMailRepository() *MailRepository {
	return &MailRepository{
		database.DbInit(),
	}
}

func (r *MailRepository) Save(mail dto.Mail) error {
	query, args, err := sq.Insert("mails").
		Columns("recipient", "subject", "upload_id", "date_sent").
		Values(mail.Recipient, mail.UploadId, mail.DateSent).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}
	_, err = r.Pool.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}

	return nil
}

func (r *MailRepository) Delete(id int) error {
	query, args, err := sq.Delete("mails").
		Where(sq.Eq{"maid_id": id}).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}
	_, err = r.Pool.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}
	return nil
}

func (r *MailRepository) DeleteAllOfRecipient(email string) error {
	query, args, err := sq.Delete("mails").
		Where(sq.Eq{"recipient": email}).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}

	_, err = r.Pool.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}
	return nil
}

func (r *MailRepository) GetByRecipient(email string) ([]dto.Mail, error) {
	var mails []dto.Mail

	query, args, err := sq.Select("*").From("mails").
		Where(sq.Eq{"recipient": email}).ToSql()
	if err != nil {
		return []dto.Mail{}, multierr.Append(lerrors.ErrInQuery, err)
	}

	rows, err := r.Pool.Query(context.Background(), query, args)
	if err != nil {
		return []dto.Mail{}, multierr.Append(lerrors.ErrNoRecord, err)
	}
	defer rows.Close()
	index := 0
	for rows.Next() {
		curr := mails[index]
		err = rows.Scan(curr.Recipient, &curr.Subject,
			&curr.UploadId, &curr.DateSent)
		if err != nil {
			return []dto.Mail{}, multierr.Append(lerrors.ErrScanningQuery, err)
		}
		index++
	}

	return mails, nil
}
