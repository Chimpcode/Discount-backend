package storage

import (
	"github.com/minio/minio-go" // Import Minio library.
	"log"
	"io"
	"io/ioutil"
	"fmt"
)


const AccessKey  = "CIUZR4SP1N4ZQA7N35FV"
const SecretKey  = "jUpBEdg2wR0zgkykbTqWK5PBpL4PzcsjZ6ladVW3"
const EndPoint = "127.0.0.1:9000"
const UseSSL = false
const Location = "sa-east-1"

const ImagesBucket = "images"

var minioClient *minio.Client

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

func UploadImage(reader io.Reader, name string) error {
	fmt.Println("-----Minio->", reader)

	n, err := minioClient.PutObject(ImagesBucket, name, reader, "image/png")
	fmt.Println("-----Minio->", err)

	if err != nil {
		return err
	}


	log.Println("Upload", n, "bytes of", name)
	return nil
}

func GetImage(name string) ([]byte, error) {
	object, err := minioClient.GetObject(ImagesBucket, name)
	if err != nil {
		return []byte{}, err
	}
	data, err := ioutil.ReadAll(object)
	return data, err
}

