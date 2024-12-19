package instance

import "github.com/minio/minio-go"

func NewMinio(endpoint string, ak string, sk string) (*minio.Client, error) {
	return minio.New(
		endpoint,
		ak,
		sk,
		false,
	)
}
