package queue

import (
	"log"
	"sync"
	"time"
)

type Sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	C   map[string]*Entry
	Rmx *sync.RWMutex
}

type Entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.Rmx.Lock()
	defer m.Rmx.Unlock()
	item, ok := m.C[key]
	if !ok {
		m.C[key] = &Entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.Rmx.RLock()
	if e, ok := m.C[key]; ok && e.isExist {
		m.Rmx.RUnlock()
		return e.value
	} else if !ok {
		m.Rmx.RUnlock()
		m.Rmx.Lock()
		e = &Entry{ch: make(chan struct{}), isExist: false}
		m.C[key] = e
		m.Rmx.Unlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	} else {
		m.Rmx.RUnlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	}
}

