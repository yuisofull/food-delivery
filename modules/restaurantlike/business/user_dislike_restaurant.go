package restaurantlikebiz

import (
	"context"
	restaurantlikemodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/model"
)

type UserDisLikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userDislikeRestaurantBiz struct {
	store UserDisLikeRestaurantStore
}

func NewUserDislikeRestaurantBiz(
	store UserDisLikeRestaurantStore,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store: store,
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

	return nil
}
