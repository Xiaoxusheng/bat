package message

import (
	"bat/config"
	"bat/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type ChatBat struct {
	Name     string
	Version  float64
	Time     string
	Upgrader websocket.Upgrader
}

//var upgrader = websocket.Upgrader{
//	HandshakeTimeout:  0,
//	ReadBufferSize:    1024,
//	WriteBufferSize:   1024,
//	WriteBufferPool:   &sync.Pool{},
//	Subprotocols:      nil,
//	Error:             nil,
//	CheckOrigin:       nil,
//	EnableCompression: false,
//}

// SendMessage 发送消息
func (c *ChatBat) SendMessage() {
	//TODO implement me
	//panic("implement me")

	ch, h, err := utils.CreateClient()
	if err != nil {
		log.Fatalf("Dial error: %v", err) // 使用 Fatalf 方便调试，打印错误后退出
	}
	fmt.Println(h.Request.URL.String())
	for {
		select {
		case message := <-config.SenChan:
			fmt.Println("收到消息", message)
			err = ch.WriteJSON(message)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}

}

// ReadMessage  接收消息
func (c *ChatBat) ReadMessage(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	//panic("implement me")

	conn, err := c.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//	循环读取消息
	for {
		var mes = new(config.MessageJson)
		err := conn.ReadJSON(mes)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(mes, err)
		if mes.MessageType == "private" {
			config.SenChan <- config.WsMessage{
				Action: "send_private_msg",
				Params: struct {
					UserID  int64  `json:"user_id"`
					Message string `json:"message"`
				}{
					UserID:  mes.UserId,
					Message: "你好",
				},
			}
		}

		//
		//config.SenChan <- config.SendMessage{
		//	UserId:      mes.UserId,
		//	Message:     "qwe",
		//	MessageType: mes.MessageType,
		//	AutoEscape:  false,
		//}
	}
}

func (c *ChatBat) Start() {
	log.Printf("%v版本%v 启动时间:[%v]", c.Name, c.Version, c.Time)
	http.HandleFunc("/", c.ReadMessage)
	go c.SendMessage()
	err := http.ListenAndServe(":"+strconv.Itoa(5700), nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func NewChatBat(version float64, readBufferSize, writeBufferSize int) *ChatBat {
	return &ChatBat{
		Name:    "ChatBat",
		Version: version,
		Time:    time.Now().Format(time.DateTime),
		Upgrader: websocket.Upgrader{
			HandshakeTimeout:  0,
			ReadBufferSize:    readBufferSize,
			WriteBufferSize:   writeBufferSize,
			WriteBufferPool:   new(sync.Pool),
			Subprotocols:      nil,
			Error:             nil,
			CheckOrigin:       nil,
			EnableCompression: false,
		},
	}
}
