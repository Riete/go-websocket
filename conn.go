package go_websocket

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Conn struct {
	conn *websocket.Conn
}

func (c *Conn) Conn() *websocket.Conn {
	return c.conn
}

func (c *Conn) WritePing() error {
	return c.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second))
}

func (c *Conn) WritePong() error {
	return c.conn.WriteControl(websocket.PongMessage, nil, time.Now().Add(time.Second))
}

func (c *Conn) WriteClose(code int, text string) error {
	return c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(code, text), time.Now().Add(time.Second))
}

func (c *Conn) WriteMessage(data []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Conn) WriteBinary(data []byte) error {
	return c.conn.WriteMessage(websocket.BinaryMessage, data)
}

func (c *Conn) WriteJson(v interface{}) error {
	return c.conn.WriteJSON(v)
}

func (c *Conn) ReadMessage() (int, []byte, error) {
	return c.conn.ReadMessage()
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func NewServer(w http.ResponseWriter, r *http.Request, h http.Header, options ...UpgraderOption) (*Conn, error) {
	upgrader := websocket.Upgrader{}
	for _, option := range options {
		option(&upgrader)
	}
	conn, err := upgrader.Upgrade(w, r, h)
	return &Conn{conn: conn}, err
}

// NewClient scheme is "ws" or "wss"
func NewClient(scheme, addr, path string, h http.Header) (*Conn, error) {
	u := url.URL{Scheme: scheme, Host: addr, Path: path}
	conn, r, err := websocket.DefaultDialer.Dial(u.String(), h)
	if err != nil {
		b, er := io.ReadAll(r.Body)
		if er != nil {
			_ = r.Body.Close()
		}
		return nil, fmt.Errorf("%s: %s %s", err.Error(), r.Status, string(b))
	}
	return &Conn{conn: conn}, nil
}
