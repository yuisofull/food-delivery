package restaurantlikestorage

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	restaurantlikemodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", data.UserId, data.RestaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
