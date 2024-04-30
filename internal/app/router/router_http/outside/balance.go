package outside

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mini_game_balance/internal/app/router/router_http/request"
	"mini_game_balance/internal/app/service/balance"
	"mini_game_balance/internal/pkg/response"
)

func GetServer(c *gin.Context) {
	var req request.GetServerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		zap.L().Error("RegisterServer", zap.Error(err))
		response.Fail(c)
		return
	}
	info, err := balance.GetServer(&req)
	var resp request.GetServerResp
	resp.ServerInfo = info
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(resp, c)
	}
}
