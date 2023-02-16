package repository

import "mime/multipart"

type UploadFileInterface interface {
	UploadFile(file *multipart.FileHeader, name string) (string, error)
	ReplaceFile(file string, newFile *multipart.FileHeader) (string, error)
	DeleteFile(file string) error
}
