package login

import (
	"auth-service/domain/entity"
	"auth-service/infra/config/database"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginService struct {
	conn *pgxpool.Pool
}

func NewLoginService() *LoginService {
	db, _ := database.DbInit()
	return &LoginService{db}
}

func (u *LoginService) SaveLogin(ctx context.Context, login entity.Login) error {
	query := "INSERT INTO login (email, password) VALUES (?,?);"
	_, err := u.conn.Exec(ctx, query,
		login.Email, login.Password)
	if err != nil {
		return fmt.Errorf("error in SaveLogin query, %v", err)
	}
	return nil
}

func (u *LoginService) GetLoginByEmailAndPassword(ctx context.Context, email string, password string) (entity.Login, error) {
	query := "SELECT * FROM login WHERE email=? AND password=?;"
	row, err := u.conn.Query(ctx, query,
		email, password)
	if err != nil {
		return entity.Login{}, fmt.Errorf("error in GetLoginByEmailAndPassword query, %v", err)
	}
	var result entity.Login
	err = row.Scan(&result)
	if err != nil {
		return entity.Login{}, fmt.Errorf("error scanning to json, %v", err)
	}
	return result, nil
}
