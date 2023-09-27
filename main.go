package main

import (
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr"  gorm:"column:addr"`
}

func (Restaurant) TableName() string { return "restaurants" }

type UpdateRestaurant struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr"  gorm:"column:addr"`
}

func (UpdateRestaurant) TableName() string { return Restaurant{}.TableName() }

func main() {
	dsn := "food_delivery:123456@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db)
	//newRestaurant := Restaurant{Name: "Land", Addr: "HCM"}
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//log.Println("New id:", newRestaurant.Id)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//POST /v1/restaurants
	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")
	restaurants.POST("", ginrestaurant.CreateRestaurant(db))

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
	restaurants.GET("", func(c *gin.Context) {
		var data []Restaurant

		pagingData := struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}{}

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}
		if pagingData.Limit <= 0 {
			pagingData.Limit = 5
		}
		if err := db.Offset((pagingData.Page - 1) * pagingData.Limit).
			Limit(pagingData.Limit).
			Order("id desc").
			Find(&data).Error; err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
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

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))
	if err := r.Run(); err != nil {
		log.Println(err)
	}
	//var myRestaurant Restaurant
	//if err := db.Where("id = ?", 3).First(&myRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//log.Println(myRestaurant)
	//
	//myRestaurant.Name = "cali"
	//if err := db.Where("id = ?", 3).Updates(&myRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//log.Println(myRestaurant)
	//
	//newName := ""
	//updateData := UpdateRestaurant{Name: &newName}
	//if err := db.Where("id = ?", 3).Updates(&updateData).Error; err != nil {
	//	log.Println(err)
	//}
	//log.Println(myRestaurant)
	//
	//if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	//	log.Println(err)
	//}
}
