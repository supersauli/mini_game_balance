package mynet

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"time"
)

type ClientRequest interface {
	ziface.IRequest
}
type ClientHandle func(request ClientRequest)

type ClientConfig struct {
	IP   string
	Port int
}

type Client struct {
	cli ziface.IClient
}

func NewClient(c *ClientConfig) *Client {
	client := &Client{}
	client.cli = znet.NewClient(c.IP, c.Port)
	if client.cli == nil {
		panic("errror")
	}
	client.cli.SetOnConnStop(func(conn ziface.IConnection) {
		for {
			newCli := znet.NewClient(c.IP, c.Port)
			if newCli.Conn().IsAlive() {
				client.cli = newCli
				break
			}
			time.Sleep(time.Second)

		}
	})
	return client
}

type handleBase struct {
	handle ClientHandle
}

func (h *handleBase) PreHandle(request ziface.IRequest) { //Hook method before processing conn business(在处理conn业务之前的钩子方法)

}

func (h *handleBase) Handle(request ziface.IRequest) { //Method for processing conn business(处理conn业务的方法)
	h.handle(request)
}

func (h *handleBase) PostHandle(request ziface.IRequest) {

}

func (c *Client) AddRouter(msgID uint32, handle ClientHandle) {

	handleBase := &handleBase{}
	handleBase.handle = handle

	c.cli.AddRouter(msgID, handleBase)
}

func (c *Client) Run() {
	c.cli.Start()
}
