package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"persia_atlas/server/websocket"
)

type Server struct {
	DB      *gorm.DB
	Router  *gin.Engine
	WsHub   *websocket.WsHub
	RedisDB *redis.Client
}

func (server *Server) Initialize() {
	server.Router = gin.Default()
	server.connectDatabase()
	server.migrateDatabase()
	server.addMiddlewares()
	server.setupRoutes()
}

func (server *Server) Run(addr string) {
	err := server.Router.Run(addr)
	if err != nil {
		log.Fatalln("error in running router:", err)
		return
	}
}
