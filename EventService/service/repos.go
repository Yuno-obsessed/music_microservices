package service

import (
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/entity"
	"github.com/Yuno-obsessed/music_microservices/EventService/domain/repository"
	"github.com/Yuno-obsessed/music_microservices/EventService/service/event"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
