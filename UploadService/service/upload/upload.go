package upload

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"
)

type UploadService struct {
	client     *minio.Client
	bucketName string
}

func NewUploadService(bucket string) *UploadService {
	ctx := context.Background()
	client, err := minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false,
	})
	if err != nil {

	}
}
