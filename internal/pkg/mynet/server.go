package mynet

import (
	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type ServerHandle interface {
	ziface.IRequest
}

type ServerConfig struct {
	//tcp4 or other
	IPVersion string
	// IP version (e.g. "tcp4") - 服务绑定的IP地址
	IP string
	// IP address the server is bound to (服务绑定的端口)
	Port int
	// 服务绑定的websocket 端口 (Websocket port the server is bound to)
	WsPort int
	// 服务绑定的kcp 端口 (kcp port the server is bound to)
	KcpPort int
}

type Server struct {
	s ziface.IServer
}

func NewServer(c *ServerConfig) *Server {
	t := &Server{}
	t.s = znet.NewServer(func(s *znet.Server) {
		s.IP = c.IP
		s.Port = c.Port
		s.IPVersion = c.IPVersion
		s.WsPort = c.WsPort
		s.KcpPort = c.KcpPort
		s.RouterSlicesMode = true
	})
	zconf.GlobalObject.RouterSlicesMode = true
	return t
}

func (s *Server) AddRoute(msgID uint32, fn func(request ServerHandle)) {
	s.s.AddRouterSlices(msgID, func(router ziface.IRequest) { fn(router) })
}

func (s *Server) Run() {
	s.s.Serve()
}
