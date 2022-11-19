package server

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"persia_atlas/server/boot"
	"persia_atlas/server/websocket"
	"persia_atlas/server/websocket/commands"
)

func init() {
	boot.LoadEnvironmentVariables()
}

func Run() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	
	wsCommands := commands.GetWsCommands()
	wsHub := websocket.NewWsHub(wsCommands)
	go wsHub.Run()

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	fmt.Printf("Listening to port %s", apiPort)

	server := Server{
		WsHub:   wsHub,
		RedisDB: rdb,
	}
	server.Initialize()
	server.Run(apiPort)
}
