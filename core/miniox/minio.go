package miniox

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/yuuki798/TimerMe3/config"
	"log"
)

var MinioClient *minio.Client

func MinioInit() (client *minio.Client, wrong bool) {
	endpoint, accessKeyID, secretAccessKey, useSSL := MinioProfile()
	if endpoint == "" || accessKeyID == "" || secretAccessKey == "" {
		log.Println("MinIO client configuration not set")
		return nil, true
	}

	//log.Println("MinIO client configuration set")

	// 初始化MinIO客户端
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Printf("Error initializing MinIO client: %v", err)
		return nil, true
	}
	log.Println("MinIO client initialized successfully")
	return minioClient, false
}

func MinioProfile() (string, string, string, bool) {
	minioConfig := config.GetConfig().Minio
	// MinIO配置
	endpoint := minioConfig.Endpoint
	accessKeyID := minioConfig.AccessKeyID
	secretAccessKey := minioConfig.SecretAccessKey
	useSSL := minioConfig.UseSSL
	return endpoint, accessKeyID, secretAccessKey, useSSL
}
