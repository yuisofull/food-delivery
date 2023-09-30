package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) ListRestaurantWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var data []restaurantmodel.Restaurant

	db := s.db.Where("status in (1)")
	if f := filter; f != nil {
		if f.OwnerID > 0 {
			db = db.Where("owner_id = ?", f.OwnerID)
		}
	}
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
