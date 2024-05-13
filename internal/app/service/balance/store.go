package balance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"mini_game_balance/configs"
	"mini_game_balance/global"
	event2 "mini_game_balance/internal/pkg/event"
	"mini_game_balance/internal/pkg/utils"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ServerTypeScoreKey = "server_type_score"
	ServerTypeInfoKey  = "server_type_info"
)

type ServerDescribe struct {
	UpdateTime int64  `json:"update_time"`
	ServeInfo  string `json:"serve_info"`
}

// var valueStore map[string]map[string]string = make(map[string]map[string]string)
var rwLock sync.RWMutex

func GetRedisKeyByServerTypeInfo(key int) string {
	return fmt.Sprintf("%s-%d", ServerTypeInfoKey, key)
}
func GetServerTypeInfoByRedisKey(key string) (int, error) {
	keyArr := strings.Split(key, "-")
	if len(keyArr) != 2 {
		return 0, errors.New("key error")
	}
	v, err := strconv.Atoi(keyArr[1])
	return v, err
}

func GetRedisKeyByServerTypeScore(key int) string {
	return fmt.Sprintf("%s-%d", ServerTypeScoreKey, key)
}

func SetServerInfo(ctx context.Context, serverType int, uid string, score float64, serverInfo string) error {
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
		// 添加到服务器列表
		key := GetRedisKeyByServerTypeInfo(serverType)
		err := rCli.HSet(context.Background(), key, map[string]string{uid: string(describeStr)}).Err()
		if err != nil {
			log.Error("redis err", zap.Error(err))
			return err
		}
	}

	{
		// 添加到排行榜
		key := GetRedisKeyByServerTypeScore(serverType)
		err := rCli.ZAdd(context.Background(), key, &redis.Z{Score: score, Member: uid}).Err()
		if err != nil {
			log.Error("redis err", zap.Error(err))
			return err
		}
	}

	return nil
}

func GetServerInfoByServerType(serverType int) (serverInfo string, err error) {
	rCli := global.RedisClient
	// 读取服务器排名
	scoreKey := GetRedisKeyByServerTypeScore(serverType)
	ranking, err := rCli.ZRangeWithScores(context.Background(), scoreKey, 0, 5).Result()
	if err != nil {
		zap.L().Error("fail ", zap.Error(err))
		return
	}
	timeNow := time.Now().Unix()
	infoKey := GetRedisKeyByServerTypeInfo(serverType)

	var getServerInfo = func(uid string) string {
		result := rCli.HGet(context.Background(), infoKey, uid)
		if result.Err() != nil {
			return ""
		}

		describe := ServerDescribe{}
		if err := json.Unmarshal([]byte(result.Val()), &describe); err != nil {
			zap.L().Error("redis err", zap.Error(err))
			return ""
		}
		// TODO del
		if describe.UpdateTime+int64(configs.ServerConfig.Balance.MaxTimeOutSec) < timeNow {
			return ""
			//rwLock.Lock()
			//defer rwLock.Unlock()
			//// 这个 服务器已经离线 移除
			//err = rCli.ZRem(context.Background(), scoreKey, uid).Err()
			//if err != nil {
			//	zap.L().Error("redis del ", zap.Error(err))
			//}
		}
		return describe.ServeInfo
	}

	// 先随机然后遍历
	if len(ranking) == 0 {
		return "", errors.New("can not found server")
	}

	for i := 0; i < 5; i++ {
		randVal := utils.RandBetween(0, len(ranking))
		uid := ranking[randVal].Member.(string)
		if serverInfo := getServerInfo(uid); serverInfo != "" {
			return serverInfo, nil
		}
	}

	for _, item := range ranking {
		uid := item.Member.(string)
		serverInfo := getServerInfo(uid)
		if serverInfo != "" {
			return serverInfo, nil
		}
	}

	return "", errors.New("can not found server")
}

// TODO tick 移除离线服务器
func RemoveOfflineServerTick() {
	rCli := global.RedisClient
	eventTime := event2.NewTicketEvent(time.Duration(configs.ServerConfig.Balance.CheckTimeOutSec) * time.Second)
	eventTime.Execute(func() bool {
		timeNow := time.Now().Unix()
		ctx := context.Background()
		rwLock.Lock()
		defer rwLock.Unlock()

		var cursor uint64 = 0
		for {

			var keys []string
			var err error
			keys, cursor, err = rCli.Scan(ctx, cursor, "server_type_info*", 0).Result() // 使用通配符匹配键
			if err != nil {
				break
			}
			// 处理匹配的键
			for _, key := range keys {
				serverType, _ := GetServerTypeInfoByRedisKey(key)
				scoreKey := GetRedisKeyByServerTypeScore(serverType)

				val, err := rCli.HGetAll(ctx, key).Result()
				if err != nil {
					continue
				}
				var removeUid []string
				for uid, v := range val {
					describe := ServerDescribe{}
					if err := json.Unmarshal([]byte(v), &describe); err != nil {
						continue
					}

					if describe.UpdateTime+180 < timeNow {
						removeUid = append(removeUid, uid)
					}

				}
				rCli.ZRem(context.Background(), scoreKey, removeUid)
			}

			if cursor == 0 {
				break
			}
		}
		return true
	})
}
