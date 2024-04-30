package request

type RegisterServerReq struct {
	ServerType int    `json:"server_type"` //服务器类型
	UID        string `json:"uid"`         // 服务唯一id
	SortKey    int    `json:"sort_key"`
	ServerInfo string `json:"server_info"` // 服务器信息
}

type RegisterServerResp struct {
	Msg string
}

type GetServerReq struct {
	ServerType int `json:"server_type"`
}
type GetServerResp struct {
	ServerInfo string `json:"server_info"`
}
