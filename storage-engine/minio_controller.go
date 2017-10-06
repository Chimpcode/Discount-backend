package storage

import (
	"github.com/minio/minio-go"
	"log"
	"io"
	"io/ioutil"
	"../global"
)


/*
const AccessKey  = "CIUZR4SP1N4ZQA7N35FV"
const SecretKey  = "jUpBEdg2wR0zgkykbTqWK5PBpL4PzcsjZ6ladVW3"
const EndPoint = "10.100.107.38:9000"

const UseSSL = false
const Location = "sa-east-1"
*/


const ImagesBucket = "images"

var MinioClient *minio.Client

func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func InitStorage(conf *global.CupointConf) {
	var err error
	MinioClient, err = minio.New(conf.MinioEndpoint, conf.MinioAk, conf.MinioSk, conf.MinioUseSsl)
	checkError(err)
	err = MinioClient.MakeBucket(ImagesBucket, conf.MinioLocation)
	if err != nil {
		exists, err := MinioClient.BucketExists(ImagesBucket)
		if err == nil && exists {
			log.Printf("We already own %s\n", ImagesBucket)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s bucket\n", ImagesBucket)

}

func UploadImage(reader io.Reader, name string) error {
	log.Println("-----Minio->", reader)

	n, err := MinioClient.PutObject(ImagesBucket, name, reader, -1, &minio.PutObjectOptions{ContentType: "image/png"})

	log.Println("-----Minio->", err)

	if err != nil {
		return err
	}

	log.Println("Upload", n, "bytes of", name)
	return nil
}

func GetImage(name string) ([]byte, error) {
	object, err := MinioClient.GetObject(ImagesBucket, name)
	if err != nil {
		return []byte{}, err
	}
	data, err := ioutil.ReadAll(object)
	return data, err
}

