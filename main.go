package main

import (
	"context"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	"github.com/yuisofull/food-delivery-app-with-go/component/uploadprovider"
	"github.com/yuisofull/food-delivery-app-with-go/middleware"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub/localpubsub"
	"github.com/yuisofull/food-delivery-app-with-go/subscriber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
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

	//  AMAZON S3
	file, err := os.Open("s3_access_keys.csv")
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
	secretKey := data[1][5]
	//s3BucketName := os.Getenv("S3BucketName")
	//s3Region := os.Getenv("S3Region")
	//s3APIKey := os.Getenv("S3APIKey")
	//s3SecretKey := os.Getenv("S3SecretKey")
	//s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	ps := localpubsub.NewPubsub()
	appCtx := appctx.NewAppContext(db, s3Provider, secretKey, ps)

	// setup subscribers
	subscriber.Setup(appCtx, context.Background())

	//  GOOGLE CLOUD
	// Using file
	//storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsFile("key.json"))
	//
	// Using env var
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

	setupRoute(appCtx, v1)
	setupAdminRoute(appCtx, v1)

	if err := r.Run(); err != nil {
		log.Println(err)
	}
}
