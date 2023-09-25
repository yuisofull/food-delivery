package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
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

	//POST /restaurants
	v1 := r.Group("/v1")
	v1.POST("/restaurants", func(c *gin.Context) {
		var data Restaurant
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Create(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	r.Run()
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
