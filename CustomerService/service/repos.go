package service

import (
	"log"

	"github.com/Yuno-obsessed/music_microservices/CustomerService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/CustomerService/service/customer"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"
	"github.com/jinzhu/gorm"
)

type Repositories struct {
	Customer customer.CustomerService
	Db       *gorm.DB
}

func NewRepositories() *Repositories {
	db, _ := gorm.Open("postgres", database.DbDns())
	log.Printf("NewRepositories, %v\n", database.DbDns())
	return &Repositories{
		Db:       db,
		Customer: customer.NewCustomerService(db),
	}
}

func (r *Repositories) Migrate() error {
	return r.Db.AutoMigrate(&entity.Customer{}).Error
}

func (r *Repositories) Close() error {
	return r.Db.Close()
}
