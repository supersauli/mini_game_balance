package main

import (
	"go.uber.org/zap"
	"mini_game_balance/configs"
	"mini_game_balance/global"
	"mini_game_balance/internal/app/router/router_http"
	"mini_game_balance/internal/pkg/db/redisDB"
	"mini_game_balance/internal/pkg/http_server"
	"mini_game_balance/internal/pkg/mylog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	configs.Init()

	// 初始化日志
	mylog.Init()

	// 初始化数据库
	redisCli, err := redisDB.NewRedis(&configs.ServerConfig.Redis)
	if err != nil {
		zap.L().Panic("init redis fail ,redis config:", zap.Any("config", configs.ServerConfig.Redis), zap.Error(err))
	} else {
		global.RedisClient = redisCli
	}

	//mysqlDB, err := sqlDB.NewMysql(&configs.ServerConfig.Mysql)
	//if err != nil {
	//	zap.L().Panic("init mysql fail ,mysql config:", zap.Any("config", configs.ServerConfig.Mysql), zap.Error(err))
	//} else {
	//	global.MySql = mysqlDB
	//}

	//model.Init()
	//tencent_storage.Init()
	//go http_server.RunOutsideServer()
	//go http_server.RunInsideServer()
	{
		httpServer := http_server.NewHttpSerer(configs.ServerConfig.System.InsideHttpAddr, router_http.NewInSideRouter())
		go httpServer.RunServer()
	}
	{
		httpServer := http_server.NewHttpSerer(configs.ServerConfig.System.OutsideHttpAddr, router_http.NewOutsideRouter())
		go httpServer.RunServer()
	}

	//web_socket.InitWebRouter()
	//go http_server.RunWebServer(
	//	configs.ServerConfig.System.WebHttpAddr,
	//	global.GinRouter,
	//	global.WebSocketRouter,
	//	global.ConnectManager)

	// 创建一个 channel，用于捕获系统信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 等待系统信号
	<-signals

	// 等待任务完成
	time.Sleep(3 * time.Second)
}
