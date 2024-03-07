package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	"github.com/yuisofull/food-delivery-app-with-go/memcache"
	"github.com/yuisofull/food-delivery-app-with-go/middleware"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/transport/ginuser"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/userstore"
)

func setupAdminRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	userStore := userstore.NewSQLStore(appCtx.GetMyDBConnection())
	userCachingStore := memcache.NewUserCaching(memcache.NewCaching(), userStore)

	admin := v1.Group("/admin", middleware.RequireAuth(appCtx, userCachingStore),
		middleware.CheckRole(appCtx, "admin"))
	{
		admin.GET("/profile", ginuser.GetProfile(appCtx))
	}
}
