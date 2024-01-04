package uploadprovider

import (
	"context"
	"fmt"
	"github.com/yuisofull/food-delivery-app-with-go/common"
)

type UploadProvider interface {
	fmt.Stringer
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
