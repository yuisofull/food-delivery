package appctx

import (
	"github.com/yuisofull/food-delivery-app-with-go/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMyDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadProvider,
	}
}

func (ctx *appCtx) GetMyDBConnection() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}
