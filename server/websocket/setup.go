package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"net/http"
	"persia_atlas/server/middlewares"
	"persia_atlas/server/models"
	"time"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to writeToWs a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer.
	AllUsersRoom   = "ALL_USERS"
	UserDirectRoom = "ROOM_"
)

type WsHandler struct {
	DB      *gorm.DB
	RedisDB *redis.Client
	wsHub   *WsHub
}

type SocketMsg struct {
	Command string `json:"command"`
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("request Sec-Websocket-Key:", r.Header.Get("Sec-Websocket-Key"))
		// check request url!
		return true
	},
}

// Client is a middleman between the websocket connection and the Hub.
type Client struct {
	ws       *websocket.Conn
	sendChan chan []byte
	user     models.User
	rooms    []string
}

type BroadcastMsg struct {
	Room string
	Data []byte
}

type WsRequest struct {
	Command int    `json:"command"`
	Payload any    `json:"payload"`
	ReqKey  string `json:"req_key"`
}

// writeToWs writes a message with the given message type and payload.
func (client *Client) writeToWs(msgType int, payload []byte) error {
	client.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return client.ws.WriteMessage(msgType, payload)
}

// readPump pumps messages from the websocket connection to the Hub.
func (client *Client) readPump(h *WsHandler, user models.User) {
	defer func() {
		h.wsHub.unregister <- client
		client.ws.Close()
	}()
	client.ws.SetReadLimit(maxMessageSize)
	client.ws.SetReadDeadline(time.Now().Add(pongWait))
	client.ws.SetPongHandler(func(string) error {
		client.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		msgType, msg, err := client.ws.ReadMessage()
		fmt.Println("readPump | msgType:", msgType, "msg:", string(msg))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Fatalf("readPump error: %v \n", err)
			}
			break
		}

		var request WsRequest
		if err := json.Unmarshal(msg, &request); err != nil {
			fmt.Println("ERROR IN PARSING REQUEST JSON!")
			msg := createClientError(request, user, "error in parsing your json")
			h.wsHub.broadcast <- msg
		} else {
			fmt.Println("unmarshal result:", request, "req key:", request.ReqKey)
			cmdHandler, ok := h.wsHub.commands[request.Command]
			if ok {
				response := cmdHandler(h, user, request)
				for _, msg := range response {
					fmt.Println("res msg:", msg.Room)
					h.wsHub.broadcast <- msg
				}
			} else {
				fmt.Println("INVALID COMMAND")
				msg := createClientError(request, user, "INVALID COMMAND")
				h.wsHub.broadcast <- msg
			}
		}
	}
}

// writePump pumps messages from the Hub to the websocket connection.
func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.ws.Close()
	}()
	for {
		select {
		case message, ok := <-client.sendChan:
			fmt.Println("writePump | message | message:", string(message))
			if !ok {
				err := client.writeToWs(websocket.CloseMessage, []byte{})
				if err != nil {
					fmt.Println("ERROR writing to ws in !ok:", err.Error())
				}
				return
			}
			err := client.writeToWs(websocket.TextMessage, message)
			if err != nil {
				fmt.Println("ERROR writing message to ws:", err.Error())
				return
			}
		case <-ticker.C:
			if err := client.writeToWs(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func WsEndpoint(h *WsHandler, w http.ResponseWriter, r *http.Request, user models.User) {
	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("could not upgrade websocket connection:", err.Error())
		return
	}

	fmt.Println("ws user:", user.Mobile)
	var rooms []string
	rooms = append(rooms, GetUserDirectRoom(user.ID))
	rooms = append(rooms, AllUsersRoom)

	client := Client{
		sendChan: make(chan []byte, 256),
		ws:       wsConn,
		user:     user,
		rooms:    rooms,
	}

	h.wsHub.register <- &client

	go client.readPump(h, user)
	go client.writePump()
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB, hub *WsHub, rdb *redis.Client) {
	h := &WsHandler{
		DB:      db,
		wsHub:   hub,
		RedisDB: rdb,
	}
	r.GET("/ws", middlewares.WsAuth(db), func(c *gin.Context) {
		user, _ := c.Get("user")
		WsEndpoint(h, c.Writer, c.Request, user.(models.User))
	})
}
