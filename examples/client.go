package main

import (
	"github.com/gorilla/websocket"
	ws "github.com/riete/go-websocket"
)

func main() {
	c, _ := ws.NewClient("ws", "127.0.0.1:8080", "echo", nil)
	c.WriteMessage([]byte("aaa"))
	c.WriteClose(websocket.CloseInternalServerErr, "bbb")
}
