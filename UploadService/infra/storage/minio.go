package storage

import (
	"context"
	"os"

	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	"github.com/Yuno-obsessed/music_microservices/ProjectLibrary/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

func NewClient() *minio.Client {
	l := logger.NewLogger()
	client, err := minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false,
	})
	if err != nil {
		l.Error("error creating minio client", zap.Error(err))
		return nil
	}
	if err = CreateBucketsIfNotExist(client); err != nil {
		l.Error("error creating buckets", zap.Error(err))
		return nil
	}

	return client
}

func CreateBucketIfNotExist(client *minio.Client, bucket consts.BucketName) error {
	exists, _ := client.BucketExists(context.Background(), string(bucket))
	if !exists {
		err := client.MakeBucket(context.Background(),
			string(bucket), minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateBucketsIfNotExist(client *minio.Client) error {
	buckets := consts.GetBucketNames()
	for i := 0; i < len(buckets); i++ {
		exists, _ := client.BucketExists(context.Background(), string(buckets[i]))
		if !exists {
			location := "us-east-1"
			err := client.MakeBucket(context.Background(),
				string(buckets[i]), minio.MakeBucketOptions{Region: location})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
