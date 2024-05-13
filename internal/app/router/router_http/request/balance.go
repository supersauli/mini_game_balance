package request

type RegisterServerReq struct {
	ServerType int                    `json:"server_type"` //服务器类型
	UID        string                 `json:"uid"`         // 服务唯一id
	ScoreMap   map[string]interface{} `json:"score_map"`
	ServerInfo string                 `json:"server_info"` // 服务器信息
}

type RegisterServerResp struct {
	ScoreResult float64 `json:"score_result"`
}

type GetServerReq struct {
	ServerType int `json:"server_type"`
}
type GetServerResp struct {
	ServerInfo string `json:"server_info"`
}
