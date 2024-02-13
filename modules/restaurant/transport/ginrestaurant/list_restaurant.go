package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	restaurantbusiness "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/business"
	restaurantmodel "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/model"
	restaurantrepo "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/repository"
	restaurantstorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/storage"
	restaurantlikestorage "github.com/yuisofull/food-delivery-app-with-go/modules/restaurantlike/store"
	"net/http"
)

func ListRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMyDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		uid, err := common.FromBase58(filter.FakeOwnerID)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.OwnerID = int(uid.GetLocalID())

		filter.Status = []int{1}

		store := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbusiness.NewListRestaurantBusiness(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
