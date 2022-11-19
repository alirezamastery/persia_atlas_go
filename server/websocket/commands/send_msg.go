package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"persia_atlas/server/models"
	"persia_atlas/server/websocket"
)

var ctx = context.Background()

type FetchData struct {
	RobotIsOn    bool `json:"robot_is_on"`
	RobotRunning bool `json:"robot_running"`
}

func SendMessageCommand(
	h *websocket.WsHandler,
	user models.User,
	request websocket.WsRequest,
) websocket.WsResponse {

	fmt.Println("send cmd | payload:", request)
	robotOn, err := h.RedisDB.Get(ctx, "robot_is_on").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		robotOn = "0"
	} else if err != nil {
		fmt.Println("error in redis get:", err.Error())
		panic(err)
	}

	robotRunning, err := h.RedisDB.Get(ctx, "robot_is_on").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		robotRunning = "0"
	} else if err != nil {
		fmt.Println("error in redis get:", err.Error())
		panic(err)
	}

	data := FetchData{
		RobotIsOn:    robotOn == "1",
		RobotRunning: robotRunning == "1",
	}
	resForMe := websocket.WsMessage{
		Type:   "fetch",
		ReqKey: request.ReqKey,
		Data:   data,
	}
	resForMeB, err := json.Marshal(resForMe)
	if err != nil {
		fmt.Println("error in marshaling fetch response:", err.Error())
		panic(err)
	}

	response := websocket.WsResponse{}
	response = append(response, websocket.BroadcastMsg{
		Room: websocket.GetUserDirectRoom(user.ID),
		Data: resForMeB,
	})

	return response
}
