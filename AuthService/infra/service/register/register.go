package register

import (
	"auth-service/domain/entity"
	"auth-service/infra/config/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RegisterService struct {
	conn *pgxpool.Pool
}

func NewRegisterService() *RegisterService {
	db, _ := database.DbInit()
	return &RegisterService{db}
}

func (r *RegisterService) SaveRegister(ctx context.Context, register entity.Register) error {
	query := "INSERT INTO register (uuid, username, email, password, country) VALUES (?,?,?,?);"
	_, err := r.conn.Exec(ctx, query,
		register.UUID, register.Username, register.Email,
		register.Password, register.Country)

	if err != nil {
		return fmt.Errorf("error in query SaveRegister, %v", err)
	}
	return nil
}

func (r *RegisterService) GetRegisterByUuid(ctx context.Context, registerUuid uuid.UUID) (entity.Register, error) {
	query := "SELECT * FROM register WHERE uuid=?;"
	row, err := r.conn.Query(ctx, query, registerUuid)
	if err != nil {
		return entity.Register{}, fmt.Errorf("error in GetRegisterByUuid query, %v", err)
	}
	var result entity.Register
	err = row.Scan(&result)
	if err != nil {
		return entity.Register{}, fmt.Errorf("error scanning to json, %v", err)
	}
	return result, nil
}

func (r *RegisterService) GetRegisterByEmailAndPassword(ctx context.Context, email string, password string) (entity.Register, error) {
	query := "SELECT * FROM register WHERE email=? AND password=?;"
	row, err := r.conn.Query(ctx, query,
		email, password)
	if err != nil {
		return entity.Register{}, fmt.Errorf("error in GetRegisterByEmailAndPassword query, %v", err)
	}
	var result entity.Register
	err = row.Scan(&result)
	if err != nil {
		return entity.Register{}, fmt.Errorf("error scanning to json, %v", err)
	}
	return result, nil
}

func (r *RegisterService) IsRegistered(ctx context.Context, email string, password string) bool {
	query := "SELECT * FROM register WHERE email=? AND password=?;"
	row, err := r.conn.Query(ctx, query,
		email, password)
	if err != nil {
		return false
	}
	var result entity.Register
	err = row.Scan(&result)
	if err != nil {
		return false
	}
	return true
}
