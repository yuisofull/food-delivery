package restaurantlikebiz

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantlikemodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/model"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	ps    pubsub.Pubsub
}

func NewUsersLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	ps pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		ps:    ps,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Send message
	if err := biz.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}
	return nil
}
