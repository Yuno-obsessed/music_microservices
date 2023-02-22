package customer

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Yuno-obsessed/music_microservices/CustomerService/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerService struct {
	db *pgxpool.Pool
}

func NewCustomerService(db *pgxpool.Pool) CustomerService {
	return CustomerService{db}
}

func (cs *CustomerService) GetById(id string) (entity.Customer, error) {
	var customer entity.Customer

	query, args, err := squirrel.Select("username", "email", "age", "country").
		From("customers").ToSql()
	if err != nil {
		return entity.Customer{}, fmt.Errorf("no such customer found in database, %v", err)
	}

	return customer, nil
}

func (cs *CustomerService) GetByUsername(username string) (entity.Customer, error) {
	var customer entity.Customer

	err := cs.db.Table("customers").Debug().
		Where("username=?", username).Take(&customer).Error
	if err != nil {
		return entity.Customer{}, fmt.Errorf("no such customer found in database, %v", err)
	}

	return customer, nil
}

func (cs *CustomerService) GetByEmailAndPassword(email string, password string) (entity.Customer, error) {
	var customer entity.Customer

	err := cs.db.Table("customers").Debug().
		Where("email=? and password=?", email, password).Take(&customer).Error
	if err != nil {
		return entity.Customer{}, fmt.Errorf("no such customer fount in database, %v", err)
	}

	return customer, nil
}

// add limit and offset
func (cs *CustomerService) GetAll() ([]entity.Customer, error) {
	var customers []entity.Customer

	err := cs.db.Table("customers").Debug().
		Limit(10).Order("desc").
		Take(&customers).Error
	if err != nil {
		return nil, fmt.Errorf("no such custimess were found, %v", err)
	}
	return customers, nil
}

func (cs *CustomerService) Create(customer entity.Customer) error {
	err := cs.db.Table("customers").Debug().
		Create(&customer).Error

	return err
}

func (cs *CustomerService) Update(customer entity.Customer) error {
	err := cs.db.Table("customers").Debug().
		Save(&customer).Error

	return err
}

func (cs *CustomerService) Delete(id string) error {
	err := cs.db.Table("customers").Debug().
		Where("customer_id=?", id).
		Delete(id).Error

	return err
}
