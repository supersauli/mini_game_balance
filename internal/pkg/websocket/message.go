package websocket

type Message struct {
	data []byte
}

func (m *Message) GetData() []byte {
	return m.data
}
