package restaurantbusiness

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
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
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (business *deleteRestaurantBusiness) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := business.store.FindRestaurantWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}
	if err := business.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
