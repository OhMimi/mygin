package test

import (
	"mygin/handler"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()

	handler.ProcessWs(c)
}

func TestWebSocketExample(t *testing.T) {

	// 建立一個測試server
	s := httptest.NewServer(http.HandlerFunc(echo))
	defer s.Close()

	// trans http://127.0.0.1 => ws://127.0.0.1
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// 嘗試連線
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	// 連線成功，發送訊息並且接收後驗證
	if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		t.Fatalf("%v", err)
	}

	_, message, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("%v", err)
	}

	result := string(message)
	if result != "got it : hello" {
		t.Fatal(result)
	}
}
