package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	restaurantbusiness "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/business"
	restaurantstorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/storage"
	"net/http"
)

func DeleteRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMyDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbusiness.NewDeleteRestaurantBusiness(store, requester)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(true))
	}
}
