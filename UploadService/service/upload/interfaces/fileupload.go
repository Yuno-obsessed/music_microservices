package interfaces

import (
	"mime/multipart"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
)

type UploadFile interface {
	UploadFile(file *multipart.FileHeader, name string, bucket consts.BucketName) (string, error)
	ReplaceFile(file string, newFile *multipart.FileHeader, bucket consts.BucketName) (string, error)
	DeleteFile(file string, bucket consts.BucketName) error
}
