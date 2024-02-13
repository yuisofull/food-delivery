package ginrstlike

import (
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	restaurantlikebiz "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/business"
	restaurantlikemodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/model"
	restaurantlikestorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/store"
	"net/http"
)

// ListUsersLikeRestaurant : GET /v1/restaurants/:id/liked-users
func ListUsersLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		//myArr := []string{}
		//
		//fmt.Println(myArr[0])

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMyDBConnection())
		biz := restaurantlikebiz.NewListUsersLikeRestaurantBiz(store)

		users, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range users {
			users[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))
	}
}
