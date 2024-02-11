package restaurantbusiness

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantmodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/model"
	"log"
)

type ListRestaurantStore interface {
	ListRestaurantWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBusiness struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

func NewListRestaurantBusiness(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantBusiness {
	return &listRestaurantBusiness{
		store:     store,
		likeStore: likeStore,
	}
}

func (business *listRestaurantBusiness) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	res, err := business.store.ListRestaurantWithCondition(ctx, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(res))

	for i := range ids {
		ids[i] = res[i].Id
	}

	likeMap, err := business.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println(err)
		return res, nil
	}

	for i := range res {
		res[i].LikedCount = likeMap[res[i].Id]
	}
	return res, nil
}
