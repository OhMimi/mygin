package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsPing(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	ProcessWs(ws)
}

func ProcessWs(ws *websocket.Conn) {
	for {
		// 讀取ws Socket傳來的訊息
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		// 如果message 是 "ping"
		if string(message) == "ping" {
			// 回應 "pong"
			message = []byte("pong")
		} else {
			// 如果是其他內容，就回應文字訊息類型，內容就是回聲
			message = []byte(fmt.Sprintf("got it : %s", string(message)))
		}

		// 寫入websocket
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
