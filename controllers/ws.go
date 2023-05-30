package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WebSocketHandler(c *gin.Context) {
	// 获取WebSocket连接
	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		return
	}

	defer ws.Close()
	// 处理WebSocket消息
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println("messageType:", messageType)
		fmt.Println("p:", string(p))
		str := "From Server " + string(p)
		// 输出WebSocket消息内容
		if err = ws.WriteMessage(messageType, []byte(str)); err != nil {
			break
		}

	}

}
