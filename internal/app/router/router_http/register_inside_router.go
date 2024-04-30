package router_http

import (
	"github.com/gin-gonic/gin"
	"mini_game_balance/internal/app/router/router_http/inside"
)

// InitInSideRouter 管理后台路由
func NewInSideRouter() *gin.Engine {
	router := gin.Default()
	baseRouter := router.Group("balance")
	registerInsideServer(baseRouter)
	return router
}

func registerInsideServer(baseRouter *gin.RouterGroup) {
	baseRouter.POST("/register_server", inside.RegisterServer)
}
