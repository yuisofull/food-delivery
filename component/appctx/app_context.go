package appctx

import (
	"github.com/yuisofull/food-delivery-app-with-go/component/uploadprovider"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMyDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider, secretKey string, ps pubsub.Pubsub) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadProvider,
		secretKey:      secretKey,
		ps:             ps,
	}
}

func (ctx *appCtx) GetMyDBConnection() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}
func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}
func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.ps
}
