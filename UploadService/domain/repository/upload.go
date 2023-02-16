package repository

import "github.com/Yuno-obsessed/music_microservices/UploadService/domain/entity"

type UploadInterface interface {
	GetByName(name string) (entity.Upload, error)
	GetByEntity(entity string) (entity.Upload, error)
	SaveUpload(event entity.Upload) error
	UpdateUpload(name string) error
	DeleteUpload(id string) error
}
