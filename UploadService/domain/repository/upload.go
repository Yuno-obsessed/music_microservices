package repository

import (
	"context"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"

	"github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/Yuno-obsessed/music_microservices/UploadRepository/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UploadRepository struct {
	Db     *pgxpool.Pool
	Logger logger.CustomLogger
}

func NewUploadRepository() *UploadRepository {
	return &UploadRepository{
		Db:     database.DbInit(),
		Logger: logger.NewLogger(),
	}
}

func (s *UploadRepository) GetByName(name string) (entity.Upload, error) {
	var upload entity.Upload

	query, args, err := squirrel.Select("*").From("uploads").
		Where(squirrel.Eq{"upload_name": name}).ToSql()
	if err != nil {
		s.Logger.Error("error building query", zap.Error(err))
		return entity.Upload{}, err
	}

	row := s.Db.QueryRow(context.Background(), query, args)
	err = row.Scan(&upload)
	if err != nil {
		s.Logger.Error("error scanning", zap.Error(err))
		return entity.Upload{}, err
	}

	return upload, nil
}

func (s *UploadRepository) GetByEntity(uentity string) ([]entity.Upload, error) {
	var upload []entity.Upload

	query, args, err := squirrel.Select("*").From("uploads").
		Where(squirrel.Eq{"upload_entity": uentity}).ToSql()
	if err != nil {
		s.Logger.Error("error building query", zap.Error(err))
		return []entity.Upload{}, err
	}

	rows, err := s.Db.Query(context.Background(), query, args)
	if err != nil {
		s.Logger.Error("error querying", zap.Error(err))
		return []entity.Upload{}, err
	}
	defer rows.Close()
	index := 0
	for rows.Next() {
		curr := upload[index]
		err = rows.Scan(&curr.UploadId, &curr.UserId,
			&curr.Uentity, &curr.Uentity)
		if err != nil {
			s.Logger.Error("error scanning", zap.Error(err))
			return []entity.Upload{}, err
		}
		index++
	}

	return upload, nil
}

func (s *UploadRepository) SaveUpload(event entity.Upload) error {
	query, args, err := squirrel.Insert("uploads").
		Columns("user_id", "upload_name", "upload_entity").
		Values(event.UserId, event.Name, event.Uentity).ToSql()
	if err != nil {
		s.Logger.Error("error building query", zap.Error(err))
		return err
	}

	_, err = s.Db.Exec(context.Background(), query, args)
	if err != nil {
		s.Logger.Error("error executing query", zap.Error(err))
		return err
	}

	return nil
}

func (s *UploadRepository) UpdateUpload(oldname, name string) error {
	query, args, err := squirrel.Update("uploads").
		Set("upload_name", name).ToSql()
	if err != nil {
		s.Logger.Error("error building query", zap.Error(err))
		return err
	}

	_, err = s.Db.Exec(context.Background(), query, args)
	if err != nil {
		s.Logger.Error("error executing query", zap.Error(err))
		return err
	}

	return nil
}

func (s *UploadRepository) DeleteUpload(id string) error {
	query, args, err := squirrel.Delete("uploads").
		Where(squirrel.Eq{"upload_id": id}).ToSql()
	if err != nil {
		s.Logger.Error("error building query", zap.Error(err))
		return err
	}

	_, err = s.Db.Exec(context.Background(), query, args)
	if err != nil {
		s.Logger.Error("error executing query", zap.Error(err))
	}

	return nil
}
