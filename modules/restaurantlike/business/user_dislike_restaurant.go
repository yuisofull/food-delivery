package restaurantlikebiz

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/component/asyncjob"
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

	// Side effects
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId)

	})

	if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
		log.Println(err)
	}
	return nil
}
