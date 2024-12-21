package session

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() SessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
		rwlock:     sync.RWMutex{}, //可以省略
	}
}

// 初始化
func (m *MemorySessionMgr) Init(addr string, option ...string) (err error) {
	return
}

// 创建session
func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	//go get github.com/satori/go.uuid
	//用uuid生成sessionId
	id := uuid.NewV4()
	//转换成字符串
	sessionId := id.String()
	//创建session
	session = NewMemorySession(sessionId)
	m.sessionMap[sessionId] = session
	return
}

// 获取session
func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	session, ok := m.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exist")
		return
	}
	return
}
