package subscriber

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	restaurantstorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/storage"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//	store := restaurantstorage.NewSQLStore(appCtx.GetMyDBConnection())
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := (msg.Data()).(HasRestaurantId)
//			err := store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//			if err != nil {
//				log.Println(err)
//			}
//		}
//	}()
//}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMyDBConnection())
			likeData := (msg.Data()).(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PushNotificationAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			likeData := (msg.Data()).(HasRestaurantId)
			log.Printf("Push notification when user likes restaurant %d \n", likeData.GetRestaurantId())
			return nil
		},
	}
}

func EmitRealtimeAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Realtime emit after user likes restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			likeData := (msg.Data()).(HasRestaurantId)
			return appCtx.GetRealtimeEngine().EmitToUser(likeData.GetUserId(), string(msg.Channel()), likeData)
		},
	}
}
