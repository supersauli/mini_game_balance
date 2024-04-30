package websocket

import "mini_game_balance/internal/proto/gen/proto_logic"

// type MsgIDType int32
// type MsgIDType proto_logic.MsgId
type Router struct {
	router map[proto_logic.MsgId]HandleFunc
}
type HandleFunc func(*Request)

func NewRouter() *Router {
	m := &Router{
		router: make(map[proto_logic.MsgId]HandleFunc),
	}

	return m
}

func (r *Router) AddRouter(msgID proto_logic.MsgId, h HandleFunc) {
	r.router[msgID] = h
}

func (r *Router) Handle(msgID proto_logic.MsgId, req *Request) {
	if v, ok := r.router[msgID]; ok {
		v(req)
	}
}
