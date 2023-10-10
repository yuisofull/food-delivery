package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMyDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(true))
	}
}
