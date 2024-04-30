package router_http

import (
	"github.com/gin-gonic/gin"
	"mini_game_balance/internal/app/router/router_http/outside"
)

// InitInSideRouter 管理后台路由
func NewOutsideRouter() *gin.Engine {
	router := gin.Default()
	baseRouter := router.Group("balance")
	registerOutsideBalanceServer(baseRouter)
	return router
}

func registerOutsideBalanceServer(baseRouter *gin.RouterGroup) {
	baseRouter.POST("/get_server", outside.GetServer)
}
