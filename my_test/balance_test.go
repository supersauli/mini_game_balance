package my

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mini_game_balance/internal/app/router/router_http/request"
	"net/http"
	"testing"
)

func TestBalanceRegister(t *testing.T) {
	for i := 1; i < 10; i++ {
		req := &request.RegisterServerReq{
			ScoreMap: map[string]interface{}{
				"user_num": 100 * i,
				"cpu":      2 * i,
			},
			ServerInfo: fmt.Sprintf("test_%d", i),
			ServerType: 1,
			UID:        fmt.Sprintf("test_%d", i),
		}
		reqJs, _ := json.Marshal(req)

		http.Post("http://127.0.0.1:20001/balance/register_server", "application/json", bytes.NewReader(reqJs))
	}

	//if score, err := balance.RegisterServer(req); err != nil {
	//	t.Error(err)
	//} else {
	//	t.Log(score)
	//}

}
func TestBalanceGetServer(t *testing.T) {
	req := &request.GetServerReq{
		ServerType: 1,
	}
	reqJs, _ := json.Marshal(req)

	resp, _ := http.Post("http://127.0.0.1:20000/balance/get_server", "application/json", bytes.NewReader(reqJs))
	defer resp.Body.Close()

	var respJS = &request.GetServerResp{}
	_ = json.NewDecoder(resp.Body).Decode(&respJS)
	t.Log(respJS)

}
