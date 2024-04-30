package websocket

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"mini_game_balance/internal/pkg/event"
	"mini_game_balance/internal/proto/gen/proto_logic"
	"sync"
)

type Connect struct {
	conn           *websocket.Conn
	connectManager *ConnectManager
	UUID           string
	propertyLock   sync.RWMutex
	property       map[string]interface{}
	router         *Router
	isValid        bool
	singleTask     *event.SingleTask
}

func NewConnect(conn *websocket.Conn, cm *ConnectManager, msgRouter *Router) *Connect {
	return &Connect{
		conn:           conn,
		connectManager: cm,
		property:       make(map[string]interface{}),
		router:         msgRouter,
		isValid:        true,
		singleTask:     event.NewSingleTask(),
	}
}

func (c *Connect) GetUID() string {
	return c.UUID
}

func (c *Connect) Close() {
	c.conn.Close()
	c.isValid = false
	zap.L().Info("SendToJson connect is close", zap.Any("connect", c))
}
func (c *Connect) SendToJson(data interface{}) {
	if !c.isValid {
		zap.L().Error("SendToJson connect is close", zap.Any("data", data))
		return
	}
	err := c.WriteJson(data)
	if err != nil {
		c.connectManager.SafeCloseConnect(c.UUID)
	}
}

func (c *Connect) SendToOtherJson(uid string, data interface{}) {
	if !c.isValid {
		zap.L().Error("SendToOtherJson connect is close", zap.Any("uid", uid), zap.Any("data", data))
		return
	}

	c.connectManager.SendToJson(uid, data)
}

func (c *Connect) SendBroadcastUIDJson(uid []string, data interface{}) {
	if !c.isValid {
		zap.L().Error("SendBroadcastUIDJson connect is close", zap.Any("uid", uid), zap.Any("data", data))
		return
	}

	c.connectManager.SendBroadcastUIDJson(uid, data)
}

func (c *Connect) WriteJson(data interface{}) error {
	if !c.isValid {
		zap.L().Error("WriteJson connect is close", zap.Any("data", data))
		return errors.New("error connect is close")
	}

	return c.conn.WriteJSON(data)
}
func (c *Connect) SendProtoBufMsg(msgID proto_logic.MsgId, message proto.Message) {
	if !c.isValid {
		zap.L().Error("SendProtoBufMsg connect is close", zap.Int("msgID", int(msgID)))
		return
	}

	data, err := proto.Marshal(message)
	if err != nil {
		zap.L().Error("Login fail", zap.Error(err))
		return
	}

	baseMsg := proto_logic.BaseMsg{}
	baseMsg.MsgId = msgID
	baseMsg.Data = data
	marshal, err := proto.Marshal(&baseMsg)
	if err != nil {
		zap.L().Error("marshal fail ", zap.String("uuid", c.UUID), zap.Error(err))
		return
	}

	marshalPack := MarshalPack(marshal)
	err = c.conn.WriteMessage(websocket.BinaryMessage, marshalPack)
	if err != nil {
		zap.L().Error("send msg fail ", zap.String("uuid", c.UUID), zap.Error(err))
		return
	}
}

func (c *Connect) WriteMessage(messageType int, data []byte) error {
	return c.conn.WriteMessage(messageType, data)
}

func (c *Connect) Run() {
	c.readLoop()
}
func (c *Connect) readLoop() {
	var messagePack MessagePack

	for {
		messageType, data, err := c.conn.ReadMessage()
		_ = messageType
		if err != nil {
			//fmt.Println("Failed to read message:", err)
			zap.L().Error("read error close ", zap.String("uuid", c.UUID), zap.Error(err))
			break
		}

		for {
			message, next := messagePack.UnmarshalPack(data)
			if message == nil {
				break
			}
			baseMsg := proto_logic.BaseMsg{}
			err = proto.Unmarshal(message, &baseMsg)
			if err != nil {
				zap.L().Error("parse fail ", zap.String("uuid", c.UUID), zap.Error(err))
				break
			}

			req := &Request{
				conn: c,
				msg:  &Message{data: baseMsg.Data},
			}

			c.singleTask.Run(func() {
				c.router.Handle(baseMsg.MsgId, req)
			})

			if !next {
				break
			}
		}
	}
	if !c.connectManager.SafeCloseConnect(c.UUID) {
		c.Close()
	}
}

func (c *Connect) DelProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.property, key)
}

func (c *Connect) SetSingleTask(task *event.SingleTask) {
	c.singleTask = task
}

func (c *Connect) GetSingleTask() *event.SingleTask {
	return c.singleTask
}

func (c *Connect) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = value
}

func (c *Connect) GetProperty(key string) (ok bool, val interface{}) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if v, ok := c.property[key]; ok {
		return true, v
	}
	return false, nil
}

func (c *Connect) GetPropertyInt64(key string) (ok bool, val int64) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if v, ok := c.property[key]; ok {
		if v2, ok := v.(int64); ok {
			return true, v2
		}
	}
	return false, 0
}

func (c *Connect) GetPropertyString(key string) (ok bool, val string) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if v, ok := c.property[key]; ok {
		if v2, ok := v.(string); ok {
			return true, v2
		}
	}
	return false, ""
}

func (c *Connect) GetPropertyInt(key string) (ok bool, val int) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if v, ok := c.property[key]; ok {
		if v2, ok := v.(int); ok {
			return true, v2
		}
	}
	return false, 0
}
