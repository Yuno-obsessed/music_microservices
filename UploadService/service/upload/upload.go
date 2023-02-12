package upload

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/lerrors"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type UploadService struct {
	client     *minio.Client
	bucketName string
	logger     logger.CustomLogger
}

func NewUploadService(bucket string) *UploadService {
	ctx := context.Background()
	logger := logger.NewLogger()
	client, err := minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false,
	})
	if err != nil {
		logger.Error("Failed to create a new minio client, " + err.Error())
		return nil
	}
	bucketName := "microservices." + bucket
	location := "us-east-1"
	exists, _ := client.BucketExists(ctx, bucketName)
	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			logger.Error("Failed to create a bucket " + bucketName + err.Error())
			return nil
		}
		logger.Info("Bucket " + bucketName + " was created successfully")
	}
	return &UploadService{client, bucketName, logger}
}

type UploadServiceInterface interface {
	UploadFile(file *multipart.FileHeader, name string) (string, error)
	ReplaceFile(file string, newFile *multipart.FileHeader) (string, error)
	DeleteFile(file string) error
}

func (u UploadService) UploadFile(file *multipart.FileHeader, name string) (string, error) {
	ctx := context.Background()
	f, err := file.Open()
	if err != nil {
		u.logger.Error(lerrors.ErrFile.Error() + err.Error())
		return "", err
	}
	defer f.Close()

	if file.Size > int64(40960000) {
		u.logger.Error(lerrors.ErrFile.Error() + "file is too big to be an avatar")
		return "", fmt.Errorf("the file is too big")
	}

	buffer := make([]byte, file.Size)
	_, err = f.Read(buffer)
	if err != nil {
		u.logger.Error(lerrors.ErrFile.Error() + "unable to read a file")
		return "", err
	}
	fileType := http.DetectContentType(buffer)
	if !strings.HasPrefix(fileType, "image") {
		u.logger.Error(lerrors.ErrFile.Error() + "the file format is invalid")
		return "", fmt.Errorf("the file format is not valid")
	}
	fileBytes := bytes.NewReader(buffer)
	fileName := file.Filename + name
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = u.client.PutObject(ctx, u.bucketName, fileName, fileBytes, file.Size,
		minio.PutObjectOptions{ContentType: fileType, UserMetadata: userMetaData})
	if err != nil {
		u.logger.Error(lerrors.ErrFile.Error() + "error putting object in the bucker")
		return "", fmt.Errorf("error putting object in the bucket %v", err)
	}
	return fileName, nil
}

func (u UploadService) ReplaceFile(file string, newFile *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	obj, err := u.client.GetObject(ctx, u.bucketName, file, minio.GetObjectOptions{})
	if err != nil {
		u.logger.Error(lerrors.ErrFile.Error() + err.Error())
		return "", err
	}
	if obj == nil {
		u.logger.Error(lerrors.ErrFile.Error() + "no such object found in the bucket")
		return "", fmt.Errorf("no such object foung in the bucket")
	}
	err = u.DeleteFile(file)
	if err != nil {
		u.logger.Error(lerrors.ErrFile.Error() + err.Error())
		return "", err
	}
	fileName, err := u.UploadFile(newFile, file)
	if err != nil {
		u.logger.Error(lerrors.ErrFile.Error() + err.Error())
		return "", err
	}
	return fileName, nil
}

func (u UploadService) DeleteFile(file string) error {
	ctx := context.Background()
	err := u.client.RemoveObject(ctx, u.bucketName, file, minio.RemoveObjectOptions{ForceDelete: true})
	if err != nil {
		u.logger.Error(lerrors.ErrFile.Error() + err.Error())
		return err
	}
	return nil
}
