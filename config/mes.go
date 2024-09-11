package config

// MessageJson 解析收到消息
type MessageJson struct {
	SelfId      int64  `json:"self_id,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Time        int    `json:"time,omitempty"`
	MessageId   int    `json:"message_id,omitempty"`
	MessageSeq  int    `json:"message_seq,omitempty"`
	RealId      int    `json:"real_id,omitempty"`
	MessageType string `json:"message_type,omitempty"`
	Sender      struct {
		UserId   int64  `json:"user_id,omitempty"`
		Nickname string `json:"nickname,omitempty"`
		Card     string `json:"card,omitempty"`
	} `json:"sender,omitempty"`
	RawMessage string `json:"raw_message,omitempty"`
	Font       int    `json:"font,omitempty"`
	SubType    string `json:"sub_type,omitempty"`
	Message    []struct {
		Type string `json:"type,omitempty"`
		Data struct {
			Text string `json:"text,omitempty"`
		} `json:"data,omitempty"`
	} `json:"message,omitempty"`
	MessageFormat string `json:"message_format,omitempty"`
	PostType      string `json:"post_type,omitempty"`
}

// SendMessage  发送消息
type SendMessage struct {
	UserId      int64  `json:"user_id"`
	GroupId     int64  `json:"group_id"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	AutoEscape  bool   `json:"auto_escape"`
}

type WsMessage struct {
	Action string `json:"action"`
	Params struct {
		UserID  int64  `json:"user_id"`
		Message string `json:"message"`
	} `json:"params"`
}

// chan
var SenChan = make(chan WsMessage, 100)
var ReadChan = make(chan MessageJson, 100)
