package instance

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinio(endpoint string, ak string, sk string) (*minio.Client, error) {
	fmt.Printf("Creating Minio client with endpoint: %s\n", endpoint)
	client, err := minio.New(
		endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(ak, sk, ""),
			Secure: true,
		},
	)
	if err != nil {
		fmt.Printf("Error creating Minio client: %v\n", err)
		return nil, err
	}
	return client, nil
}
