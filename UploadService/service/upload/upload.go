package upload

import (
	"github.com/Yuno-obsessed/music_microservices/UploadService/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UploadService struct {
	db *pgxpool.Pool
}

func NewUploadService(db *pgxpool.Pool) UploadService {
	return UploadService{db}
}

func (us UploadService) GetByName(name string) (entity.Upload, error) {
	var upload entity.Upload

}

func (us UploadService) GetByEntity(entity string) (entity.Upload, error) {

}

func (us UploadService) SaveUpload(event entity.Upload) error {

}

func (us UploadService) UpdateUpload(name string) error {

}

func (us UploadService) DeleteUpload(id string) error {

}

