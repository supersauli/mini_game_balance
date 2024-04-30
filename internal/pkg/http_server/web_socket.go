package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mini_game_balance/internal/pkg/websocket"
	"net/http"
	"time"
)

func RunWebServer(addr string, engine *gin.Engine, router *websocket.Router, connectManager *websocket.ConnectManager) {
	address := fmt.Sprintf("%s", addr)
	server := websocket.NewServer(router, connectManager)
	//ginRouter := gin.Default()
	//global.GinRouter.Use(middleware.WebSocketJwtAuth()).GET("/ws", server.WebSockHandle)
	engine.GET("/game_logic/ws", server.WebSockHandle)
	zap.L().Info("启动 gin 路由", zap.String("地址", address))
	s := &http.Server{
		Addr:           address,
		Handler:        engine,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	zap.L().Error(s.ListenAndServe().Error())
}
