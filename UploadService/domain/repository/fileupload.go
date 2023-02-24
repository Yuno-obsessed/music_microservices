package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"go.uber.org/multierr"

	"github.com/Yuno-obsessed/music_microservices/UploadRepository/infra/storage"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/minio/minio-go/v7"
)

type FileUploadRepository struct {
	*minio.Client
}

func NewFileUploadRepository() *FileUploadRepository {
	return &FileUploadRepository{
		storage.NewClient(),
	}
}

func (r *FileUploadRepository) UploadFile(file *multipart.FileHeader, name string, bucket consts.BucketName) (string, error) {
	ctx := context.Background()
	f, err := file.Open()
	if err != nil {
		return "", multierr.Append(lerrors.ErrFile, err)
	}
	defer f.Close()

	if file.Size > int64(40960000) {
		return "", errors.New("the file is too big")
	}

	buffer := make([]byte, file.Size)
	_, err = f.Read(buffer)
	if err != nil {
		return "", multierr.Append(lerrors.ErrFile, err)
	}
	fileType := http.DetectContentType(buffer)
	if !strings.HasPrefix(fileType, "image") {
		return "", fmt.Errorf("the file format is not valid")
	}
	fileBytes := bytes.NewReader(buffer)
	fileName := file.Filename + name
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = r.Client.PutObject(ctx, string(bucket), fileName, fileBytes, file.Size,
		minio.PutObjectOptions{ContentType: fileType, UserMetadata: userMetaData})
	if err != nil {
		return "", multierr.Append(lerrors.ErrMinio, err)
	}
	return fileName, nil
}

func (r *FileUploadRepository) ReplaceFile(file string, newFile *multipart.FileHeader, bucket consts.BucketName) (string, error) {
	ctx := context.Background()
	obj, err := r.Client.GetObject(ctx, string(bucket), file, minio.GetObjectOptions{})
	if err != nil {
		return "", multierr.Append(lerrors.ErrMinio, err)
	}
	if obj == nil {
		return "", multierr.Append(lerrors.ErrNoObject, err)
	}
	err = r.DeleteFile(file, bucket)
	if err != nil {
		return "", multierr.Append(lerrors.ErrMinio, err)
	}
	fileName, err := r.UploadFile(newFile, file, bucket)
	if err != nil {
		return "", multierr.Append(lerrors.ErrMinio, err)
	}
	return fileName, nil
}

func (r *FileUploadRepository) DeleteFile(file string, bucket consts.BucketName) error {
	ctx := context.Background()
	err := r.Client.RemoveObject(ctx, string(bucket), file, minio.RemoveObjectOptions{ForceDelete: true})
	if err != nil {
		return multierr.Append(lerrors.ErrMinio, err)
	}
	return nil
}
