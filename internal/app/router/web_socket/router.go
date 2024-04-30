package web_socket

import (
	"github.com/gin-gonic/gin"
	"mini_game_balance/global"
	"mini_game_balance/internal/pkg/http_server"
	"mini_game_balance/internal/pkg/websocket"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.Use(http_server.Cors())
	return router
}
func InitWebRouter() {
	global.GinRouter = Routers()
	global.WebSocketRouter = websocket.NewRouter()
	global.ConnectManager = websocket.NewConnectManager()

	registerCommonRouter()
}
