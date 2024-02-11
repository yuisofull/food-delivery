package restaurantbusiness

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantmodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantRepo struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBusiness(repo ListRestaurantRepo) *listRestaurantRepo {
	return &listRestaurantRepo{repo: repo}
}

func (biz *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.repo.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
