package business

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(context.Context, *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBusiness(store CreateRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (business *createRestaurantBusiness) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name can not be empty")
	}

	if err := business.store.CreateRestaurant(context, data); err != nil {
		return err
	}
	return nil
}
