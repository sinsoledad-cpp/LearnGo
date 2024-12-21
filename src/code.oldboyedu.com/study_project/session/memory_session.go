package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	return &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
		//rwlock:    sync.RWMutex{}, //可删除
	}
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exist in memory session")
		return
	}
	return

}

func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}
