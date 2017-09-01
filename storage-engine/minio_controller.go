package storage

import (
	"github.com/minio/minio-go" // Import Minio library.
	"log"
)


const AccessKey  = "CIUZR4SP1N4ZQA7N35FV"
const SecretKey  = "jUpBEdg2wR0zgkykbTqWK5PBpL4PzcsjZ6ladVW3"
const EndPoint = "http://127.0.0.1:9000"
const UseSSL = true
const Location = "us-east-1"

const ImagesBucket = "images"


func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func InitStorage() {
	minioClient, err := minio.New(EndPoint, AccessKey, SecretKey, UseSSL)
	checkError(err)
	err = minioClient.MakeBucket(ImagesBucket, Location)
	if err != nil {
		exists, err := minioClient.BucketExists(ImagesBucket)
		if err == nil && exists {
			log.Printf("We already own %s\n", ImagesBucket)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", ImagesBucket)

}


