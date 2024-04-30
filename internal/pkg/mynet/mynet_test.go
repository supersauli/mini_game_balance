package mynet

import (
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	conf := &ServerConfig{
		IPVersion: "tcp",
		IP:        "127.0.0.1",
		Port:      8081,
	}
	server := NewServer(conf)
	server.AddRoute(1, func(request ServerHandle) {
		time.Sleep(time.Second * 1)
		request.GetConnection().SendMsg(1, []byte("server"))
		t.Log("server")
	})
	server.Run()
}

func TestClient(t *testing.T) {
	conf := &ClientConfig{
		IP:   "127.0.0.1",
		Port: 8081,
	}
	client := NewClient(conf)
	client.AddRouter(1, func(request ClientRequest) {
		time.Sleep(time.Second)
		request.GetConnection().SendMsg(1, []byte("client"))
		t.Log("client")
	})
	go func() {
		for {
			time.Sleep(time.Second * 2)
			err := client.cli.Conn().SendMsg(1, []byte("11"))
			if err != nil {
				t.Error(err)
			}
			t.Log("client 1")
		}
	}()
	client.Run()

	select {}

}
