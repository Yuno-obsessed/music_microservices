package repository

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/Yuno-obsessed/music_microservices/UploadRepository/infra/storage"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/minio/minio-go/v7"
)

type FileUploadRepository struct {
	Client *minio.Client
	Logger logger.CustomLogger
}

func NewFileUploadRepository() *FileUploadRepository {
	return &FileUploadRepository{
		Client: storage.NewClient(),
		Logger: logger.NewLogger(),
	}
}

func (r *FileUploadRepository) UploadFile(file *multipart.FileHeader, name string, bucket consts.BucketName) (string, error) {
	ctx := context.Background()
	f, err := file.Open()
	if err != nil {
		r.Logger.Error(lerrors.ErrFile.Error() + err.Error())
		return "", err
	}
	defer f.Close()

	if file.Size > int64(40960000) {
		r.Logger.Error(lerrors.ErrFile.Error() + "file is too big to be an avatar")
		return "", fmt.Errorf("the file is too big")
	}

	buffer := make([]byte, file.Size)
	_, err = f.Read(buffer)
	if err != nil {
		r.Logger.Error(lerrors.ErrFile.Error() + "unable to read a file")
		return "", err
	}
	fileType := http.DetectContentType(buffer)
	if !strings.HasPrefix(fileType, "image") {
		r.Logger.Error(lerrors.ErrFile.Error() + "the file format is invalid")
		return "", fmt.Errorf("the file format is not valid")
	}
	fileBytes := bytes.NewReader(buffer)
	fileName := file.Filename + name
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = r.Client.PutObject(ctx, bucket, fileName, fileBytes, file.Size,
		minio.PutObjectOptions{ContentType: fileType, UserMetadata: userMetaData})
	if err != nil {
		r.Logger.Error(lerrors.ErrFile.Error() + "error putting object in the bucker")
		return "", fmt.Errorf("error putting object in the bucket %v", err)
	}
	return fileName, nil
}

func (r *FileUploadRepository) ReplaceFile(file string, newFile *multipart.FileHeader, bucket consts.BucketName) (string, error) {
	ctx := context.Background()
	obj, err := r.Client.GetObject(ctx, bucket, file, minio.GetObjectOptions{})
	if err != nil {
		r.Logger.Error(lerrors.ErrFile.Error() + err.Error())
		return "", err
	}
	if obj == nil {
		r.Logger.Error(lerrors.ErrFile.Error() + "no such object found in the bucket")
		return "", fmt.Errorf("no such object foung in the bucket")
	}
	err = r.DeleteFile(file, bucket)
	if err != nil {
		r.Logger.Error(lerrors.ErrFile.Error() + err.Error())
		return "", err
	}
	fileName, err := r.UploadFile(newFile, file, bucket)
	if err != nil {
		r.Logger.Error(lerrors.ErrFile.Error() + err.Error())
		return "", err
	}
	return fileName, nil
}

func (r *FileUploadRepository) DeleteFile(file string, bucket consts.BucketName) error {
	ctx := context.Background()
	err := r.Client.RemoveObject(ctx, bucket, file, minio.RemoveObjectOptions{ForceDelete: true})
	if err != nil {
		r.Logger.Error(lerrors.ErrFile.Error() + err.Error())
		return err
	}
	return nil
}
