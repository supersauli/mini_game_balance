package http_server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpServer struct {
	addr   string
	router *gin.Engine
}

func NewHttpSerer(addr string, router *gin.Engine) *HttpServer {
	return &HttpServer{
		addr:   addr,
		router: router,
	}
}

func NewRoute() *gin.Engine {
	routers := gin.Default()
	routers.Use(Cors())
	return routers
}

func (s *HttpServer) RunServer() {
	address := s.addr
	zap.L().Info("启动 gin 路由", zap.String("地址", address))

	httpServer := &http.Server{
		Addr:           address,
		Handler:        s.router,
		ReadTimeout:    5 * 60 * time.Second,
		WriteTimeout:   5 * 60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	zap.L().Error(httpServer.ListenAndServe().Error())
}
