package uploadprovider

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"io"
)

type gCloudProvider struct {
	bucketName    string
	storageClient *storage.Client
	domain        string
}

func NewGCloudProvider(
	bucketName string,
	storageClient *storage.Client,
	domain string,
) *gCloudProvider {
	return &gCloudProvider{
		bucketName:    bucketName,
		storageClient: storageClient,
		domain:        domain,
	}
}

func (provider *gCloudProvider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	sw := provider.storageClient.Bucket(provider.bucketName).Object(dst).NewWriter(ctx)

	if _, err := io.Copy(sw, fileBytes); err != nil {
		return nil, err
	}

	if err := sw.Close(); err != nil {
		return nil, err
	}
	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "Google Cloud",
	}

	return img, nil
}
