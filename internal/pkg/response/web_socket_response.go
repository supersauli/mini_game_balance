package response

import "mini_game_balance/internal/pkg/websocket"

type WebSockResponse struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	MsgID string      `json:"msg_id"`
}

func WebSocketResult(msgID string, code int, data interface{}, msg string, c *websocket.Connect) {
	// 开始时间
	c.SendToJson(WebSockResponse{
		code,
		data,
		msg,
		msgID,
	})
}

func WebSocketOk(msgID string, c *websocket.Connect) {
	WebSocketResult(msgID, SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func WebSocketOkWithMessage(msgID string, message string, c *websocket.Connect) {
	WebSocketResult(msgID, SUCCESS, map[string]interface{}{}, message, c)
}

func WebSocketOkWithData(msgID string, data interface{}, c *websocket.Connect) {
	WebSocketResult(msgID, SUCCESS, data, "操作成功", c)
}

func WebSocketOkWithDetailed(msgID string, data interface{}, message string, c *websocket.Connect) {
	WebSocketResult(msgID, SUCCESS, data, message, c)
}

func WebSocketFail(msgID string, c *websocket.Connect) {
	WebSocketResult(msgID, ERROR, map[string]interface{}{}, "操作失败", c)
}

func WebSocketFailWithMessage(msgID string, message string, c *websocket.Connect) {
	WebSocketResult(msgID, ERROR, map[string]interface{}{}, message, c)
}

func WebSocketFailWithError(msgID string, err error, c *websocket.Connect) {
	//c.Set("BusinessError", err.Error())
	WebSocketResult(msgID, ERROR, map[string]interface{}{}, err.Error(), c)
}

func WebSocketFailWithDetailed(msgID string, data interface{}, message string, c *websocket.Connect) {
	WebSocketResult(msgID, ERROR, data, message, c)
}
