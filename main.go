package main

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	"github.com/yuisofull/food-delivery-app-with-go/component/uploadprovider"
	"github.com/yuisofull/food-delivery-app-with-go/middleware"
	"github.com/yuisofull/food-delivery-app-with-go/modules/restaurant/transport/ginrestaurant"
	"github.com/yuisofull/food-delivery-app-with-go/modules/upload/uploadtransport/ginupload"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/usertransport/ginuser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
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
	db = db.Debug()
	//newRestaurant := Restaurant{Name: "Land", Addr: "HCM"}
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//log.Println("New id:", newRestaurant.Id)

	//  AMAZON S3

	file, err := os.Open("S3_accessKeys.csv")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	s3BucketName := data[1][0]
	s3Region := data[1][1]
	s3Domain := data[1][2]
	s3APIKey := data[1][3]
	s3SecretKey := data[1][4]

	//s3BucketName := os.Getenv("S3BucketName")
	//s3Region := os.Getenv("S3Region")
	//s3APIKey := os.Getenv("S3APIKey")
	//s3SecretKey := os.Getenv("S3SecretKey")
	//s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appctx.NewAppContext(db, s3Provider)

	//  GOOGLE CLOUD
	//using file
	//storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsFile("key.json"))
	//
	////using env var
	////storageClient, err := storage.NewClient(context.Background(),
	////option.WithCredentialsJSON([]byte(os.Getenv("GCLOUD_STORAGE_CREDENTIAL"))))
	//
	//if err != nil {
	//	panic(err)
	//}
	//gcloudProvider := uploadprovider.NewGCloudProvider("food-deliver", storageClient, "https://storage.googleapis.com/food-deliver")
	//appCtx := appctx.NewAppContext(db, gcloudProvider)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("/static", "./static")

	//POST /v1/restaurants
	v1 := r.Group("/v1")

	//POST /v1/upload
	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))

	restaurants := v1.Group("/restaurants")
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
