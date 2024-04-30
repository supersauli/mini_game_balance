package event

import (
	"log"
	"sync/atomic"
	"time"
)

// 定时器事件

type TicketEvent struct {
	ticketTime time.Duration
	stop       atomic.Bool
}

// 创建定时器
func NewTicketEvent(t time.Duration) *TicketEvent {
	return &TicketEvent{ticketTime: t}
}

// 定时器
func (t *TicketEvent) Execute(fn func() bool) {

	// 定时器
	ticket := time.NewTicker(t.ticketTime)
	for {
		if t.stop.Load() {
			break
		}
		select {
		case _ = <-ticket.C:
			func() {
				defer func() {
					if err := recover(); err != nil {
						// 生成dump
						log.Println("TicketEvent:Execute panic err ", err)
					}
				}()
				ret := fn()
				if !ret {
					return
				}
			}()
		}
	}
}

// 停止定时器
func (t *TicketEvent) Stop() {
	t.stop.Store(true)
}
