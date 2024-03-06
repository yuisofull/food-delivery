package skio

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/yuisofull/food-delivery-app-with-go/component/tokenprovider/jwt"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/transport/skuser"
	"github.com/yuisofull/food-delivery-app-with-go/modules/user/userstore"
	"gorm.io/gorm"
	"log"
	"sync"
)

type AppContext interface {
	GetMyDBConnection() *gorm.DB
	GetSecretKey() string
}

type RealtimeEngine interface {
	//UserSockets (userId int) []AppSocket : because a user can use system (online) at the same time on many platforms such as web (2 tabs), mobile, Ipad, ...
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(ctx AppContext, engine *gin.Engine) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (engine *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}
}

func (engine *rtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	return engine.storage[userId]
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	if scks, ok := engine.storage[userId]; ok {
		return scks
	}
	return []AppSocket{}
}

func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}
func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)

	for _, s := range sockets {
		s.Emit(key, data)
	}

	return nil
}

func (engine *rtEngine) Run(appCtx AppContext, r *gin.Engine) error {
	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	if err != nil {
		return err
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected: ", s.ID(), "Ip: ", s.RemoteAddr(), s.ID())

		s.Emit("test", "world")
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {

		// Validate token
		// If false: s.Close(), and return

		// If true
		// => UserId
		// Fetch db find user by Id
		// Here: s belongs t who? {user_id}
		// We need a map[user_id][]socketio.Conn

		db := appCtx.GetMyDBConnection()
		store := userstore.NewSQLStore(db)

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
			s.Close()
			return
		}
		user.Mask(false)

		appSck := NewAppSocket(s, user)
		engine.saveAppSocket(user.Id, appSck)

		s.Emit("authenticated", user)

		//appSck.Join(user.GetRole())

		server.OnEvent("/", "UserUpdateLocation", skuser.OnUserUpdateLocation(appCtx, user))
	})

	server.OnEvent("/", "test", func(s socketio.Conn, msg interface{}) {
		log.Println("test:", msg)
	})

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
		log.Println("server receive notice:", p.Name, p.Age)

		p.Age = 33
		s.Emit("notice", p)
	})

	go func() {
		err := server.Serve()
		if err != nil {
			log.Println(err)
		}
	}()

	r.GET("socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	return nil
}
