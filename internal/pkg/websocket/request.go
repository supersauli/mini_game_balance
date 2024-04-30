package websocket

type Request struct {
	conn *Connect
	msg  *Message
}

func (r *Request) GetConn() *Connect {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}
