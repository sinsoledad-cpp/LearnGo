package session

import (
	"errors"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

type RedisSessionMgr struct {
	//redis地址
	addr string
	//密码
	password string
	//连接池
	pool *redis.Pool
	//大map
	sessionMap map[string]Session
	//读写锁
	rwlock sync.RWMutex
}

func NewRedisSessionMgr() SessionMgr {
	return &RedisSessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}

func (r *RedisSessionMgr) Init(addr string, option ...string) (err error) {
	if len(option) > 0 {
		r.password = option[0]
	}
	r.addr = addr
	r.pool = myPool(addr, r.password)
	return
}

func myPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1024,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			//认证
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		//测试连接
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//go get github.com/satori/go.uuid
	//用uuid生成sessionId
	id := uuid.NewV4()
	//转换成字符串
	sessionId := id.String()
	//创建session
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exist")
		return
	}
	return
}
