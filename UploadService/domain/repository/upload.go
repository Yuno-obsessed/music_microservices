package repository

import (
	"context"

	"github.com/Yuno-obsessed/music_microservices/UploadRepository/domain/dto"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"go.uber.org/multierr"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"

	"github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UploadRepository struct {
	*pgxpool.Pool
}

func NewUploadRepository() *UploadRepository {
	return &UploadRepository{
		database.DbInit(),
	}
}

func (s *UploadRepository) GetByName(name string) (entity.Upload, error) {
	var upload entity.Upload

	query, args, err := squirrel.Select("*").From("uploads").
		Where(squirrel.Eq{"upload_name": name}).ToSql()
	if err != nil {
		return entity.Upload{}, multierr.Append(lerrors.ErrInQuery, err)
	}

	row := s.QueryRow(context.Background(), query, args)
	err = row.Scan(&upload)
	if err != nil {
		return entity.Upload{}, multierr.Append(lerrors.ErrScanningQuery, err)
	}

	return upload, nil
}

func (s *UploadRepository) GetByEntity(uentity string) ([]entity.Upload, error) {
	var upload []entity.Upload

	query, args, err := squirrel.Select("*").From("uploads").
		Where(squirrel.Eq{"upload_entity": uentity}).ToSql()
	if err != nil {
		return []entity.Upload{}, multierr.Append(lerrors.ErrInQuery, err)
	}

	rows, err := s.Query(context.Background(), query, args)
	if err != nil {
		return []entity.Upload{}, multierr.Append(lerrors.ErrInQuery, err)
	}
	defer rows.Close()
	index := 0
	for rows.Next() {
		curr := upload[index]
		err = rows.Scan(&curr.UploadId, &curr.UserId,
			&curr.Uentity, &curr.Uentity)
		if err != nil {
			return []entity.Upload{}, multierr.Append(lerrors.ErrScanningQuery, err)
		}
		index++
	}

	return upload, nil
}

func (s *UploadRepository) SaveUpload(upload dto.UploadDto) error {
	query, args, err := squirrel.Insert("uploads").
		Columns("user_id", "upload_name", "upload_entity").
		Values(upload.UserId, upload.Name, upload.Uentity).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}

	_, err = s.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}

	return nil
}

func (s *UploadRepository) UpdateUpload(oldname, name string) error {
	query, args, err := squirrel.Update("uploads").
		Set("upload_name", name).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}

	_, err = s.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}

	return nil
}

func (s *UploadRepository) DeleteUpload(id string) error {
	query, args, err := squirrel.Delete("uploads").
		Where(squirrel.Eq{"upload_id": id}).ToSql()
	if err != nil {
		return multierr.Append(lerrors.ErrInQuery, err)
	}

	_, err = s.Exec(context.Background(), query, args)
	if err != nil {
		return multierr.Append(lerrors.ErrExecQuery, err)
	}

	return nil
}
