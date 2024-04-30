package web_socket

import (
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"mini_game_balance/global"
	"mini_game_balance/internal/pkg/websocket"
	"mini_game_balance/internal/proto/gen/proto_logic"
)

func registerCommonRouter() {
	global.WebSocketRouter.AddRouter(proto_logic.MsgId_HEARTBEAT, HeartBeat)

}

func HeartBeat(s *websocket.Request) {
	req := &proto_logic.HeartbeatRequest{}
	err := proto.Unmarshal(s.GetData(), req)
	if err != nil {
		zap.L().Error("Login fail", zap.Error(err))
		return
	}

	resp := &proto_logic.HeartbeatResponse{
		Data: req.Data,
	}

	s.GetConn().SendProtoBufMsg(proto_logic.MsgId_HEARTBEAT_RET, resp)
}
