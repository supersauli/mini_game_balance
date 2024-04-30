package grpc_server

import (
	"context"
	"log"
	"mini_game_balance/internal/proto/gen/proto_logic"
	"net"

	"google.golang.org/grpc"
)

type RpcServer struct {
	addr      string
	protoType string
}

// protoType = "tcp" or "udp ,addr = ":50051"
func NewRcpServer(addr, protoType string) *RpcServer {
	return &RpcServer{
		addr:      addr,
		protoType: protoType,
	}
}

type handle struct {
	proto_logic.UnimplementedBaseMsgCallServer
}

func (s *handle) Add(ctx context.Context, in *proto_logic.BaseMsg) (*proto_logic.BaseMsg, error) {
	//ctx := metadata.AppendToOutgoingContext(context.Background(), md)

	//metadata.FromIncomingContext(ctx)
	resp := &proto_logic.BaseMsg{
		Data: []byte("hello"),
	}
	return resp, nil
}

func (s *RpcServer) Run() {
	lis, err := net.Listen(s.protoType, s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto_logic.RegisterBaseMsgCallServer(grpcServer, &handle{})

	log.Println("gRPC server listening on port %s type %s", s.addr, s.protoType)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
