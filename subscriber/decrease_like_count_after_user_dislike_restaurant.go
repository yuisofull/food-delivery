package subscriber

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	restaurantstorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/storage"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub"
)

//func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserDislikeRestaurant)
//	store := restaurantstorage.NewSQLStore(appCtx.GetMyDBConnection())
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := (msg.Data()).(HasRestaurantId)
//			err := store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
//			if err != nil {
//				log.Println(err)
//			}
//		}
//	}()
//}

func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user dislikes restaurant",
		Hld: func(ctx context.Context, msg *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMyDBConnection())
			likeData := (msg.Data()).(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
