package main

import (
	"log"
	"net/http"

	ws "github.com/riete/go-websocket"
)

func echo(w http.ResponseWriter, r *http.Request) {
	s, _ := ws.NewServer(w, r, nil, ws.WithDisableCheckOrigin())
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
