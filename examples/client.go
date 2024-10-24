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
	c.SetPingHandler(func(s string) error {
		log.Println("recv ping from server: " + s)
		return nil
	})
	ch := c.SendHeartbeat(context.Background(), time.Second, 3*time.Second, []byte("hello I'm client"))
	go func() {
		log.Println(<-ch)
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			log.Println(c.WriteMessage([]byte("aaa")))
		}
	}()
	c.ReadMessage()
	c.WriteClose(websocket.CloseInternalServerErr, "bbb")
	c.Close()
}
