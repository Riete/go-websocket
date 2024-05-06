package main

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket"
	ws "github.com/riete/go-websocket"
)

func main() {
	c, _ := ws.NewClient(nil, "ws", "127.0.0.1:8080", "echo", nil)
	c.SetPongHandler(func(s string) error {
		log.Println("recv ping from server")
		return nil
	})
	ch := c.SetHeartbeat(context.Background(), time.Second, 3*time.Second)
	go func() {
		log.Println(<-ch)
	}()
	go c.ReadMessage()
	for {
		time.Sleep(time.Second)
		log.Println(c.WriteMessage([]byte("aaa")))
	}
	c.WriteClose(websocket.CloseInternalServerErr, "bbb")
}
