package subscriber

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
)

func Setup(appCtx appctx.AppContext, ctx context.Context) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, ctx)
	DecreaseLikeCountAfterUserDislikeRestaurant(appCtx, ctx)
	PushNotificationAfterUserLikeRestaurant(appCtx, ctx)
}
