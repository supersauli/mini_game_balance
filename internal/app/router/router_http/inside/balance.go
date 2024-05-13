package inside

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mini_game_balance/internal/app/router/router_http/request"
	"mini_game_balance/internal/app/service/balance"
	"mini_game_balance/internal/pkg/response"
)

func RegisterServer(c *gin.Context) {
	var req request.RegisterServerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		zap.L().Error("RegisterServer", zap.Error(err))
		response.Fail(c)
		return
	}
	score, err := balance.RegisterServer(&req)
	var resp request.RegisterServerResp
	resp.ScoreResult = score
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessageData("success", resp, c)
	}
}
