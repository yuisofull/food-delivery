package restaurantlikebiz

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/component/asyncjob"
	restaurantlikemodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/model"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikedCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikedCountResStore
}

func NewUsersLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	incStore IncLikedCountResStore,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:    store,
		incStore: incStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Side effects
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
		log.Println(err)
	}
	return nil
}
