package balance

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"mini_game_balance/global"
	"sync"
	"time"
)

type ServerDescribe struct {
	UpdateTime int64  `json:"update_time"`
	ServeInfo  string `json:"serve_info"`
}

// var valueStore map[string]map[string]string = make(map[string]map[string]string)
var rwLock sync.RWMutex

func GetRedisKeyByServerTypeInfo(key int) string {
	return fmt.Sprintf("server_type_info_%d", key)
}

func GetRedisKeyByServerTypeScore(key int) string {
	return fmt.Sprintf("server_type_score_%d", key)
}
func SetServerInfo(ctx context.Context, serverType int, uid string, score int, serverInfo string) error {
	rwLock.Lock()
	defer rwLock.Unlock()

	log := ctx.Value("log").(*zap.Logger)
	if log == nil {
		log = zap.L()
	}
	describe := ServerDescribe{}
	describe.ServeInfo = serverInfo
	describe.UpdateTime = time.Now().Unix()
	describeStr, _ := json.Marshal(describe)

	rCli := global.RedisClient

	{
		key := GetRedisKeyByServerTypeInfo(serverType)
		err := rCli.HSet(context.Background(), key, map[string]string{uid: string(describeStr)}).Err()
		if err != nil {
			log.Error("redis err", zap.Error(err))
			return err
		}
	}

	{
		key := GetRedisKeyByServerTypeScore(serverType)
		err := rCli.ZAdd(context.Background(), key, &redis.Z{Score: float64(score), Member: uid}).Err()
		if err != nil {
			log.Error("redis err", zap.Error(err))
			return err
		}
	}

	//{
	//	if v, ok := valueStore[key]; ok {
	//		v[uid] = serverInfo
	//	} else {
	//		valueStore[GetRedisKeyByServerType(serverType)] = map[string]string{uid: serverInfo}
	//	}
	//}

	return nil
}

func GetServerInfoByServerType(serverType int) (serverInfo string, err error) {
	rCli := global.RedisClient
	// 读取服务器排名
	scoreKey := GetRedisKeyByServerTypeScore(serverType)
	ranking, err := rCli.ZRangeWithScores(context.Background(), scoreKey, 0, 1).Result()
	if err != nil {
		zap.L().Error("fail ", zap.Error(err))
		return
	}
	timeNow := time.Now().Unix()
	infoKey := GetRedisKeyByServerTypeInfo(serverType)
	for _, item := range ranking {
		uid := item.Member.(string)
		result := rCli.HGet(context.Background(), infoKey, uid)
		if result.Err() != nil {
			continue
		}

		describe := ServerDescribe{}
		if err := json.Unmarshal([]byte(result.Val()), &describe); err != nil {
			zap.L().Error("redis err", zap.Error(err))
			continue
		}

		if describe.UpdateTime+180 < timeNow {
			func() {
				rwLock.Lock()
				defer rwLock.Unlock()
				// 这个 服务器已经离线 移除
				err = rCli.ZRem(context.Background(), scoreKey, uid).Err()
				if err != nil {
					zap.L().Error("redis del ", zap.Error(err))
				}
			}()

			continue
		}
		return describe.ServeInfo, nil
	}

	//key := GetRedisKeyByServerTypeInfo(serverType)
	//{
	//
	//	result := rCli.HGetAll(context.Background(), key)
	//	if result.Err() != nil {
	//		zap.L().Error("redis err", zap.Error(result.Err()))
	//		err = result.Err()
	//	}
	//	val = result.Val()
	//
	//}
	//{
	//	if v, ok := valueStore[key]; ok {
	//		val = v
	//	}
	//
	//}

	return
}
