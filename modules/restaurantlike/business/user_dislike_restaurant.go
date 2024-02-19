package restaurantlikebiz

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantlikemodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/model"
	"log"
)

type UserDisLikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

type DecLikedCountResStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store    UserDisLikeRestaurantStore
	decStore DecLikedCountResStore
}

func NewUserDislikeRestaurantBiz(
	store UserDisLikeRestaurantStore,
	decStore DecLikedCountResStore,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store:    store,
		decStore: decStore,
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

	go func() {
		defer common.AppRecover()
		if err := biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()
	return nil
}
