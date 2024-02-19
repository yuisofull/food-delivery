package ginrstlike

import (
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	restaurantstorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/storage"
	restaurantlikebiz "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/business"
	restaurantlikemodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/model"
	restaurantlikestorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/store"
	"net/http"
)

// UserDislikeRestaurant :
// DELETE /v1/restaurants/:id/liked-users
func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMyDBConnection())
		decStore := restaurantstorage.NewSQLStore(appCtx.GetMyDBConnection())
		biz := restaurantlikebiz.NewUserDislikeRestaurantBiz(store, decStore)

		if err = biz.DislikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(true))
	}
}
