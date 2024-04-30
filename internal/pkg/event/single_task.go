package event

import (
	"log"
	"sync"
)

type SingleTaskGroup struct {
	lock  sync.RWMutex
	group map[string]*SingleTask
}

func NewSingleTaskGroup() *SingleTaskGroup {
	return &SingleTaskGroup{
		group: make(map[string]*SingleTask),
	}
}

func (s *SingleTaskGroup) Add(key string, task *SingleTask) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.group[key]; !ok {
		s.group[key] = task
		return true
	}
	return false
}

func (s *SingleTaskGroup) Del(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.group, key)
}

func (s *SingleTaskGroup) Get(key string) *SingleTask {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.group[key]
}

// 创建一个单例任务
type SingleTask struct {
	lock sync.Mutex
}

func NewSingleTask() *SingleTask {
	return &SingleTask{}
}

func (s *SingleTask) Run(fn func()) {
	s.lock.Lock()
	defer s.lock.Unlock()
	func() {
		defer func() {
			if err := recover(); err != nil {
				// 生成dump
				log.Println("SingleTask:Run panic err ", err)
			}
		}()
		fn()
	}()
}
