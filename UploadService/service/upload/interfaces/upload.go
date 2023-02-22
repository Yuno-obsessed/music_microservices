package interfaces

import "github.com/Yuno-obsessed/music_microservices/UploadRepository/domain/entity"

type Upload interface {
	GetByName(name string) (entity.Upload, error)
	GetByEntity(uentity string) ([]entity.Upload, error)
	SaveUpload(event entity.Upload) error
	UpdateUpload(oldname, name string) error
	DeleteUpload(id string) error
}
