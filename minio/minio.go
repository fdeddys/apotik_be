package miniocon

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/minio/minio-go"
)

var (
	minioClient     *minio.Client
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
	err             error
)

func init() {
	//beego config
	endpoint = beego.AppConfig.String("minio.endpoint")
	accessKeyID = beego.AppConfig.String("minio.accessKeyID")
	secretAccessKey = beego.AppConfig.String("minio.secretAccessKey")
	useSSL, err = beego.AppConfig.Bool("minio.useSSL")

	// Initialize minio client object.
	ConnectMinio()
}

func ConnectMinio() {
	minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		fmt.Println("Error :", err)

	}
}

func GetMiniConn() *minio.Client {
	return minioClient
}
