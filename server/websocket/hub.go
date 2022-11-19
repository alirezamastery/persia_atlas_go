package websocket

import (
	"fmt"
	"persia_atlas/server/models"
	"strings"
)

type WsMessage struct {
	Type   string `json:"type"`
	Data   any    `json:"data"`
	ReqKey string `json:"req_key"`
}

type WsResponse []BroadcastMsg

type CommandHandler func(h *WsHandler, u models.User, r WsRequest) WsResponse

// WsHub maintains the set of active connections and broadcasts
// messages to the connections.
type WsHub struct {
	rooms      map[string]map[*Client]bool
	broadcast  chan BroadcastMsg
	register   chan *Client
	unregister chan *Client
	commands   map[int]CommandHandler
}

func NewWsHub(cmdMap map[int]CommandHandler) *WsHub {
	return &WsHub{
		broadcast:  make(chan BroadcastMsg),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[string]map[*Client]bool),
		commands:   cmdMap,
	}
}

func (hub *WsHub) Run() {
	fmt.Println("Run the Websocket Hub...")
	for {
		select {
		case client := <-hub.register:
			fmt.Println("Hub | register | user:", client.user.Mobile)
			for _, room := range client.rooms {
				clients := hub.rooms[room]
				if clients == nil {
					clients = make(map[*Client]bool)
					hub.rooms[room] = clients
				}
				hub.rooms[room][client] = true
			}

		case client := <-hub.unregister:
			fmt.Println("Hub | unregister | user:", client.user.Mobile)
			for _, room := range client.rooms {
				clients := hub.rooms[room]
				if clients != nil {
					if _, ok := clients[client]; ok {
						fmt.Println("Hub | unregister | close channel:", clients)
						delete(clients, client)
						if len(clients) == 0 {
							delete(hub.rooms, room)
						}
					}
				}
			}
			close(client.sendChan)

		case msg := <-hub.broadcast:
			fmt.Println(strings.Repeat("-", 100))
			fmt.Println("Hub | broadcast | room:", msg.Room)
			clients := hub.rooms[msg.Room]
			fmt.Println("Hub | broadcast | client num:", len(clients))
			for c := range clients {
				select {
				case c.sendChan <- msg.Data:
					fmt.Println("Hub | broadcast | msg.data:", string(msg.Data))
				default:
					fmt.Println("Hub | broadcast | default ----")
					close(c.sendChan)
					delete(clients, c)
					if len(clients) == 0 {
						delete(hub.rooms, msg.Room)
					}
				}
			}
			fmt.Println(strings.Repeat("-", 100))
		}
	}
}
