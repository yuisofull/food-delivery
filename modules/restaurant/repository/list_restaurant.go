package restaurantrepo

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantmodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/model"
	"log"
)

type ListRestaurantStore interface {
	ListRestaurantWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

// RestaurantLikeStore to optimize algorithm: []restaurantlikemodel.Like => map[int]int
// because when Join restaurant_likes table & restaurant = O(n^2)
// when use Map, we can reduce the algorithm = O(n)
type RestaurantLikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore RestaurantLikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore,
	likeStore RestaurantLikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		likeStore: likeStore,
	}
}

func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	res, err := repo.store.ListRestaurantWithCondition(ctx, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(res))

	for i := range ids {
		ids[i] = res[i].Id
	}

	likeMap, err := repo.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println(err)
		return res, nil
	}

	for i := range res {
		res[i].LikedCount = likeMap[res[i].Id]
	}
	return res, nil
}
