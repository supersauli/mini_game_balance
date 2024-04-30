package grpc_server

import "mini_game_balance/internal/proto/gen/proto_logic"

type MessageHandle struct {
}

func NewMessageHandle() *MessageHandle {
	return &MessageHandle{}
}

func (m *MessageHandle) RegisterMessageHandle(server *proto_logic.UnimplementedBaseMsgCallServer) {
}
