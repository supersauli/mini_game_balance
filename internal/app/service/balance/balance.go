package balance

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"mini_game_balance/internal/app/router/router_http/request"
	"mini_game_balance/internal/pkg/utils"
)

func RegisterServer(req *request.RegisterServerReq) error {
	reqUUID := utils.RanInt64()
	log := zap.L().With(zap.String("func_name", "RegisterServer"), zap.Int64("req_uuid", reqUUID))
	log.Debug("request ", zap.Any("req", req))

	ctx := context.Background()

	ctx = context.WithValue(ctx, "log", log)

	err := SetServerInfo(ctx, req.ServerType, req.UID, req.SortKey, req.ServerInfo)

	return err
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
