package miniox

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/yuuki798/TimerMe3/core/libx"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func Download(c *gin.Context, bucketName string, objectName string) {
	if bucketName == "" || objectName == "" {
		libx.Err(c, http.StatusInternalServerError, "Bucket and object parameters are required", nil)
		return
	}

	object, err := MinioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		libx.Err(c, http.StatusInternalServerError, "Could not retrieve the file", err)
		return
	}
	defer func(object *minio.Object) {
		err := object.Close()
		if err != nil {
			libx.Err(c, http.StatusInternalServerError, "Error closing the file", err)
		}
	}(object)

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", objectName))
	c.Header("Content-Type", "application/octet-stream")
	if _, err = io.Copy(c.Writer, object); err != nil {
		libx.Err(c, http.StatusInternalServerError, "Error sending the file", err)
		return
	}
	libx.Ok(c, "File sent successfully", nil)
}

func Upload(c *gin.Context, bucketName string, objectName string) {
	if bucketName == "" || objectName == "" {
		libx.Err(c, http.StatusInternalServerError, "Bucket and object parameters are required", nil)
		return
	}

	file, err := c.FormFile("uploadFile")
	if err != nil {
		libx.Err(c, http.StatusBadRequest, "Error retrieving the file", err)
		return
	}

	srcFile, err := file.Open()
	if err != nil {
		libx.Err(c, http.StatusInternalServerError, "Error opening the file", err)
		return
	}
	defer func(srcFile multipart.File) {
		err := srcFile.Close()
		if err != nil {
			libx.Err(c, http.StatusInternalServerError, "Error closing the file", err)
		}
	}(srcFile)

	_, err = MinioClient.PutObject(context.Background(), bucketName, objectName, srcFile, file.Size, minio.PutObjectOptions{})
	if err != nil {
		libx.Err(c, http.StatusInternalServerError, "Error uploading the file", err)
		return
	}
	libx.Ok(c, "File uploaded successfully", nil)
}
func DownloadToLocal(c *gin.Context, client *minio.Client, bucketName string, objectName string, filePath string) error {
	// 构建保存路径，将文件存储到当前目录的 tmp 文件夹下
	tmpFilePath := fmt.Sprintf("tmp/%s", filePath)

	// 从MinIO存储桶中下载文件
	err := client.FGetObject(c, bucketName, objectName, tmpFilePath, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("Error downloading file from MinIO: %v", err)
		libx.Err(c, http.StatusInternalServerError, "从MinIO下载文件失败", err)
		return err
	}
	log.Println("File downloaded from MinIO successfully")
	return nil
}
