package main

import (
	"context"
	"log"
	"net/http"
	"time"

	ws "github.com/riete/go-websocket"
)

func echo(w http.ResponseWriter, r *http.Request) {
	s, _ := ws.NewServer(w, r, nil, ws.WithDisableCheckOrigin())
	s.SetPingHandler(func(s string) error {
		log.Println("recv ping from client: " + s)
		return nil
	})
	ch := s.SendHeartbeat(context.Background(), time.Second, 3*time.Second, []byte("hello I'm server"))
	go func() {
		log.Println("heartbeat", <-ch)
	}()
	defer s.Close()
	for {
		mt, message, err := s.ReadMessage()
		log.Println("=====", mt, string(message), err)
		if err != nil {
			break
		}
	}
	log.Println("quit")
}

func main() {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
