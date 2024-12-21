package session

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	//设置session，可以先放到内存的map中
	//批量导入redis，提升性能
	sessionMap map[string]interface{}
	//读写锁
	rwlock sync.RWMutex
	//记录内存中的map是否被操作
	flag int
}

const (
	//内存数据没变化
	SessionFlagNone = iota
	//内存数据有变化
	SessionFlagModify
)

// 构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		sessionMap: make(map[string]interface{}, 16),
		pool:       pool,
		flag:       SessionFlagNone,
	}
	return s
}

func (r *RedisSession) Set(key string, value interface{}) (err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.RUnlock()
	//设置session
	r.sessionMap[key] = value
	//设置标记
	r.flag = SessionFlagModify
	return
}

func (r *RedisSession) Save() (err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.RUnlock()
	//如果数据没有变化，不需要保存
	if r.flag != SessionFlagModify {
		return
	}
	//内存中的session数据进行序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	//获取redis连接
	conn := r.pool.Get()
	//保存kv
	_, err = conn.Do("SET", r.sessionId, data)
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Get(key string) (result interface{}, err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.RUnlock()
	//先判断内存
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exist in redis session")
	}
	return
}

// 从redis中加载数据
func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	replay, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	//转字符串
	data, err := redis.String(replay, err)
	if err != nil {
		return
	}
	//反序列化
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return
}
func (r *RedisSession) Del(key string) (err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.RUnlock()
	//删除session
	delete(r.sessionMap, key)
	return
}
