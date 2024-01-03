package uploadprovider

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
