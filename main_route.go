package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	"github.com/yuisofull/food-delivery-app-with-go/middleware"
	"github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/transport/ginrestaurant"
	"github.com/yuisofull/food-delivery-app-with-go/modules/upload/uploadtransport/ginupload"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/transport/ginuser"
	"log"
	"net/http"
	"strconv"
)

func setupRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	db := appCtx.GetMyDBConnection()

	//POST /v1/upload
	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))

	v1.POST("/authenticate", ginuser.Login(appCtx))

	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

	//GET/v1/restaurants/:id
	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data Restaurant
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//GET/v1/restaurants/
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	//Patch
	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data UpdateRestaurant
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
}
