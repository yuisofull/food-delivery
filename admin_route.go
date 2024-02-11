package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	"github.com/yuisofull/food-delivery-app-with-go/middleware"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/transport/ginuser"
)

func setupAdminRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin", middleware.RequireAuth(appCtx), middleware.CheckRole(appCtx, "admin"))
	{
		admin.GET("/profile", ginuser.GetProfile(appCtx))
	}
}
