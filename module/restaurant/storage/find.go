package restaurantstorage

import (
	"context"
	"errors"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantmodel "github.com/yuisofull/food-delivery-app-with-go/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindRestaurantWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
