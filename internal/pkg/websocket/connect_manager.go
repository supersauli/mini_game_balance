package websocket

import (
	"errors"
	"sync"

	"go.uber.org/zap"
)

type ConnectManager struct {
	connectGroup map[string]*Connect
	lock         sync.Mutex
}

func NewConnectManager() *ConnectManager {
	return &ConnectManager{
		connectGroup: make(map[string]*Connect),
	}
}

// AddConnect 增加新的连接
func (m *ConnectManager) AddConnect(conn *Connect) error {
	defer m.lock.Unlock()
	m.lock.Lock()
	uid := conn.GetUID()
	if uid == "" {
		return errors.New("connect uuid is empty")
	}
	if v, ok := m.connectGroup[uid]; ok {
		//TODO send close msg ?
		v.Close()
	}
	m.connectGroup[uid] = conn
	return nil
}

func (m *ConnectManager) DelConnect(conn *Connect) {
	defer m.lock.Unlock()
	m.lock.Lock()
	uid := conn.GetUID()
	if v, ok := m.connectGroup[uid]; ok {
		v.Close()
	}
	delete(m.connectGroup, uid)
}

// SendToJson 给指定用户发送json消息
func (m *ConnectManager) SendToJson(uid string, data interface{}) {
	defer m.lock.Unlock()
	m.lock.Lock()
	if v, ok := m.connectGroup[uid]; ok {
		if err := v.WriteJson(data); err != nil {
			//TODO log
			m.closeConnect(uid)
		}
	}
}

// SendBroadcastUIDJson 指定用户消息
func (m *ConnectManager) SendBroadcastUIDJson(uidList []string, data interface{}) {
	//TODO unsafe
	defer m.lock.Unlock()
	m.lock.Lock()
	for _, uid := range uidList {
		if v, ok := m.connectGroup[uid]; ok {
			if err := v.WriteJson(data); err != nil {
				m.closeConnect(uid)
			}
		}
	}
}

// 获得用户连接
func (m *ConnectManager) GetConnect(uid string) *Connect {
	defer m.lock.Unlock()
	m.lock.Lock()
	if v, ok := m.connectGroup[uid]; ok {
		return v
	}
	return nil
}
func (m *ConnectManager) GetConnectByOption(option func(*Connect) bool) *Connect {
	defer m.lock.Unlock()
	m.lock.Lock()
	for _, v := range m.connectGroup {
		if option(v) {
			return v
		}
	}
	return nil
}

// closeConnect 关闭连接
func (m *ConnectManager) closeConnect(uuid string) {
	//defer m.lock.Unlock()
	//m.lock.Lock()
	if v, ok := m.connectGroup[uuid]; ok {
		v.Close()
	}
	delete(m.connectGroup, uuid)
	zap.L().Error("conn close ", zap.String("uuid", uuid))
}

// 关闭连接
func (m *ConnectManager) SafeCloseConnect(uuid string) bool {
	defer m.lock.Unlock()
	m.lock.Lock()
	zap.L().Error("conn close ", zap.String("uuid", uuid))
	if v, ok := m.connectGroup[uuid]; ok {
		v.Close()
		delete(m.connectGroup, uuid)
		return true
	}
	return false
}
