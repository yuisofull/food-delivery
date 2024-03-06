package ginuser

import (
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	"github.com/yuisofull/food-delivery-app-with-go/component/hasher"
	"github.com/yuisofull/food-delivery-app-with-go/component/tokenprovider/jwt"
	userbiz "github.com/yuisofull/food-delivery-app-with-go/modules/user/business"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/usermodel"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/userstore"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMyDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey()) //appctx.SecretKey()

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(appCtx, store, 60*60*24*30, tokenProvider, md5)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(account))
	}
}
