package skuser

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"gorm.io/gorm"
	"log"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type SmallAppContext interface {
	GetMyDBConnection() *gorm.DB
}

func OnUserUpdateLocation(appCtx SmallAppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		log.Println("User", requester.GetUserId(), "at location", location)
	}
}
