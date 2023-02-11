package service

import (
	"event-service/domain/entity"
	"event-service/domain/repository"
	"event-service/service/event"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"project_library/database"
)

type Repositories struct {
	Event repository.EventRepository
	Db    *gorm.DB
}

func NewRepositories() *Repositories {
	db, _ := gorm.Open("postgres", database.DbDns())
	return &Repositories{
		Db:    db,
		Event: event.NewEventService(db),
	}
}

func (r *Repositories) Migrate() error {
	return r.Db.AutoMigrate(&entity.Event{}).Error
}

func (r *Repositories) Close() error {
	return r.Db.Close()
}
