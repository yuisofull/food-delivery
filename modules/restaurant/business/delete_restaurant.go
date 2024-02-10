package restaurantbusiness

import (
	"context"
	"errors"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantmodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindRestaurantWithCondition(
		context.Context,
		map[string]interface{},
		...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context.Context, int) error
}

type deleteRestaurantBusiness struct {
	store     DeleteRestaurantStore
	requester common.Requester
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore, requester common.Requester) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store, requester: requester}
}

func (business *deleteRestaurantBusiness) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := business.store.FindRestaurantWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
		} else {
			return common.ErrDB(err)
		}
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}
	if oldData.UserID != business.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}
	if err := business.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
