package ginupload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	"net/http"
)

func UploadImage(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		if err = c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleNewSuccessResponse(common.Image{
			Id:        0,
			Url:       "http://localhost:8080/static/" + fileHeader.Filename,
			Width:     0,
			Height:    0,
			CloudName: "local",
			Extension: "",
		}))
	}
}
