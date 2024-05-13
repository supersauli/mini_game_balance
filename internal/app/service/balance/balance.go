package balance

import (
	"context"
	"errors"
	"github.com/dengsgo/math-engine/engine"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"mini_game_balance/configs"
	"mini_game_balance/internal/app/router/router_http/request"
	"mini_game_balance/internal/pkg/utils"
	"strings"
)

func RegisterServer(req *request.RegisterServerReq) (score float64, err error) {
	reqUUID := utils.RanInt64()
	log := zap.L().With(zap.String("func_name", "RegisterServer"), zap.Int64("req_uuid", reqUUID))
	log.Debug("request ", zap.Any("req", req))

	ctx := context.Background()

	ctx = context.WithValue(ctx, "log", log)

	formula := configs.ServerConfig.Balance.Formula
	for k, v := range req.ScoreMap {
		formula = strings.ReplaceAll(formula, k, cast.ToString(v))
	}

	score, err = engine.ParseAndExec(formula)
	if err != nil {
		log.Error("engine exec formula error ", zap.Any("req", req), zap.Error(err))
		return
	}

	err = SetServerInfo(ctx, req.ServerType, req.UID, score, req.ServerInfo)
	if err != nil {
		log.Error("set server info error ", zap.Any("req", req), zap.Error(err))
	}

	return
}

func GetServer(req *request.GetServerReq) (string, error) {
	//reqUUID := utils.RanInt64()
	//log := zap.L().With(zap.String("func_name", "GetServer"), zap.Int64("req_uuid", reqUUID))
	//log.Debug("request ", zap.Any("req", req))

	//ctx := context.Background()

	//ctx = context.WithValue(ctx, "log", log)

	info, err := GetServerInfoByServerType(req.ServerType)
	if len(info) == 0 {
		err = errors.New("server_empty")
	}

	return info, err
}
