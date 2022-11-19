package websocket

import (
	"encoding/json"
	"fmt"
	"persia_atlas/server/models"
)

func GetUserDirectRoom(userId uint) string {
	return fmt.Sprintf("%s%d", UserDirectRoom, userId)
}

func createClientError(r WsRequest, user models.User, msg string) BroadcastMsg {
	errMsg := WsMessage{
		Type:   "error",
		ReqKey: r.ReqKey,
		Data:   msg,
	}
	errMsgB, _ := json.Marshal(errMsg)
	return BroadcastMsg{
		Room: GetUserDirectRoom(user.ID),
		Data: errMsgB,
	}
}
