package event

import (
	"testing"
	"time"
)

// 测试定时事件
func TestTicketEvent(t *testing.T) {
	e := NewTicketEvent(time.Second * 2)
	go e.Execute(func() bool {
		//zap.L().Error("dd")
		t.Log("ticket event", time.Now().Format(time.TimeOnly))
		panic("test panic")
		return true
	})

	e2 := NewTicketEvent(time.Minute)
	e2.Execute(func() bool {
		t.Log("ticket event2", time.Now().Format(time.TimeOnly))
		e.Stop()
		return true
	})
}

// 测试消息事件
func TestMessageEvent(t *testing.T) {
	e := NewMessageGroup()
	e.RegisterEvent(1, -1, func(arg interface{}) {
		t.Log("message event", arg, "index -1")
	})

	e.RegisterEvent(1, 0, func(arg interface{}) {
		t.Log("message event", arg, "index 0")
	})
	e.RegisterEvent(1, 1, func(arg interface{}) {
		t.Log("message event", arg, "index 1")
	})
	e.ExecEvent(1, "this is test")
}
