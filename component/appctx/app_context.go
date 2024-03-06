package appctx

import (
	"github.com/yuisofull/food-delivery-app-with-go/component/uploadprovider"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub"
	skio "github.com/yuisofull/food-delivery-app-with-go/socketio"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMyDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
	GetPubSub() pubsub.Pubsub
	GetRealtimeEngine() skio.RealtimeEngine
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
	rtEngine       skio.RealtimeEngine
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
func (ctx *appCtx) GetRealtimeEngine() skio.RealtimeEngine {
	return ctx.rtEngine
}
func (ctx *appCtx) SetRealtimeEngine(rtEngine skio.RealtimeEngine) {
	ctx.rtEngine = rtEngine
}
