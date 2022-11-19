package commands

import "persia_atlas/server/websocket"

func GetWsCommands() map[int]websocket.CommandHandler {
	cmdMap := map[int]websocket.CommandHandler{
		1: SendMessageCommand,
	}
	return cmdMap
}
