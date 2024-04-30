package message_middleware

type MessageMiddleware struct {
	messageType int //
}

func Init(registerMessageType int) *MessageMiddleware {
	return &MessageMiddleware{
		messageType: registerMessageType,
	}
}

// 发送消息流
func (m *MessageMiddleware) SendMessage(messageType int, message interface{}) {

}

// 添加消息接收管道
func (m *MessageMiddleware) AddReceiveMessageChan(messageType int, delChan chan []byte) {

}
