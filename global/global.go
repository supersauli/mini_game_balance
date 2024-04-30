package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"mini_game_balance/internal/pkg/event"
	"mini_game_balance/internal/pkg/websocket"
)

// GinRouter 路由
var GinRouter *gin.Engine

// WebSocketRouter
var WebSocketRouter *websocket.Router

var ConnectManager *websocket.ConnectManager

var MySql *gorm.DB
var RedisClient *redis.Client

var EventGroup *event.MessageGroup = event.NewMessageGroup()

// 全局单利任务
var SingleTask = event.NewSingleTask()
