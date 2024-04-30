package websocket

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

type Server struct {
	cm     *ConnectManager
	router *Router
}

func NewServer(router *Router, cm *ConnectManager) *Server {
	return &Server{
		router: router,
		cm:     cm,
	}
}

// 定义一个全局的 upgrader 变量，用于将 HTTP 连接升级为 WebSocket 连接
var upGrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) WebSockHandle(c *gin.Context) {
	// 定义一个 HTTP 请求处理函数，该函数会在浏览器向服务器发送 WebSocket 请求时被调用
	// 将 HTTP 连接升级为 WebSocket 连接
	zap.L().Info("new conn", zap.String("ip", c.RemoteIP()))
	//var down chan int
	go func() {
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			//fmt.Println("Failed to upgrade connection:", err)
			return
		}

		conn := NewConnect(ws, s.cm, s.router)
		conn.Run()

	}()

}
