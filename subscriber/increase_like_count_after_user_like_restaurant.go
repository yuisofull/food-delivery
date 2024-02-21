package subscriber

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	restaurantstorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/storage"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
	store := restaurantstorage.NewSQLStore(appCtx.GetMyDBConnection())
	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := (msg.Data()).(HasRestaurantId)
			err := store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
			if err != nil {
				log.Println(err)
			}
		}
	}()
}

func PushNotificationAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := (msg.Data()).(HasRestaurantId)
			log.Printf("Push notification when user likes restaurant %d \n", likeData.GetRestaurantId())
		}
	}()
}
