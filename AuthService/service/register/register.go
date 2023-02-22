package register

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/AuthService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type RegisterService struct {
	db     *pgxpool.Pool
	logger logger.CustomLogger
}

func NewRegisterService(db *pgxpool.Pool) *RegisterService {
	return &RegisterService{
		db,
		logger.NewLogger(),
	}
}

func (r *RegisterService) SaveRegister(u entity.User) error {
	query, args, err := squirrel.Insert("users").
		Columns("username", "password", "age", "country", "email", "role_id").
		Values(u.Username, u.Password, u.Age, u.Country, u.Email, u.RoleId).ToSql()
	if err != nil {
		r.logger.Error("error building query in SaveRegister", zap.Error(err))
		return fmt.Errorf("error in query SaveRegister, %v", err)
	}

	_, err = r.db.Exec(context.Background(), query, args)
	if err != nil {
		r.logger.Error("error executing query in SaveRegister", zap.Error(err))
		return err
	}
	return nil
}

func (r *RegisterService) GetRegisterById(registerId int) (entity.User, error) {
	var user entity.User

	query, args, err := squirrel.Select("*").From("users").
		Where(squirrel.Eq{"user_id": registerId}).ToSql()
	if err != nil {
		r.logger.Error("error building query in GetRegisterById", zap.Error(err))
		return entity.User{}, fmt.Errorf("error in query GetRegisterById, %v", err)
	}

	row := r.db.QueryRow(context.Background(), query, args)
	err = row.Scan(&user)
	if err != nil {
		r.logger.Error("error scanning in GetRegisterById", zap.Error(err))
		return entity.User{}, err
	}

	return user, nil
}

func (r *RegisterService) GetRegisterByEmailAndPassword(email string, password string) (entity.User, error) {
	var user entity.User

	query, args, err := squirrel.Select("*").From("users").
		Where(squirrel.Eq{"email": email, "password": password}).ToSql()
	if err != nil {
		r.logger.Error("error building query in SaveRegister", zap.Error(err))
		return entity.User{}, fmt.Errorf("error in query GetRegisterByEmailAndPassword, %v", err)
	}

	row := r.db.QueryRow(context.Background(), query, args)
	err = row.Scan(&user)
	if err != nil {
		r.logger.Error("error scanning in GetRegisterById", zap.Error(err))
		return entity.User{}, err
	}

	return user, nil
}
