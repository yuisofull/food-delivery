package restaurantlikebiz

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantlikemodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/model"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub"
	"log"
)

type UserDisLikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userDislikeRestaurantBiz struct {
	store UserDisLikeRestaurantStore
	ps    pubsub.Pubsub
}

func NewUserDislikeRestaurantBiz(
	store UserDisLikeRestaurantStore,
	ps pubsub.Pubsub,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store: store,
		ps:    ps,
	}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Delete(ctx, data)
	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	// Send message
	if err := biz.ps.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}
	return nil
}
