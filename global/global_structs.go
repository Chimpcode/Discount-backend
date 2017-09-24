package global

type CupointConf struct {
	MinioEndpoint string `json:"minio_endpoint"`
	MinioAk string `json:"minio_ak"`
	MinioSk string `json:"minio_sk"`
	MinioUseSsl bool `json:"minio_use_ssl"`
	MinioLocation string `json:"minio_location"`
}
