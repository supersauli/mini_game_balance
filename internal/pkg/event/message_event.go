package event

import "sort"

// 消息事件

type MessageGroup struct {
	list map[int][]struct {
		fn    func(arg interface{})
		index int
	}
}

func NewMessageGroup() *MessageGroup {
	eventGroup := &MessageGroup{}
	eventGroup.list = make(map[int][]struct {
		fn    func(arg interface{})
		index int
	})

	return eventGroup
}

// 注册事件 事件id ，排序，执行函数
func (e *MessageGroup) RegisterEvent(eventID, index int, fn func(arg interface{})) {

	if v, ok := e.list[eventID]; ok {
		v = append(v, struct {
			fn    func(arg interface{})
			index int
		}{fn: fn, index: index})
		sort.Slice(v, func(i, j int) bool {
			return v[j].index > v[i].index
		})
		e.list[eventID] = v
	} else {
		e.list[eventID] = append(v, struct {
			fn    func(arg interface{})
			index int
		}{fn: fn, index: index})
	}
}

// 发送事件
func (e *MessageGroup) ExecEvent(eventID int, arg interface{}) {
	if v, ok := e.list[eventID]; ok {
		for _, e := range v {
			e.fn(arg)
		}
	}
}
