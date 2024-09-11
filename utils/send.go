package utils

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func CreateClient() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保在函数结束时取消上下文

	// 创建 WebSocket 连接
	dialer := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second, // 设置握手超时时间
	}

	// 发送请求头（如果需要，可以在这里添加自定义头）
	headers := http.Header{}
	// headers.Add("Authorization", "Bearer your_token") // 添加自定义头（示例）

	c, h, err := dialer.DialContext(ctx, "ws://localhost:3000/", headers)
	if err != nil {
		log.Fatalf("Dial error: %v", err) // 使用 Fatalf 方便调试，打印错误后退出
		return nil, nil, err
	}
	log.Println("Dial success client addr:", h.Request.URL.String())
	return c, h, nil
}
