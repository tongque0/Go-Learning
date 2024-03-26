package cache

import (
	"cache/internal/util"
	"log"
	"sync"
	"time"
)

type memCache struct {
	memorySizeStr     string
	maxMemorySize     int64
	currMemorySize    int64
	defaultExpireTime time.Duration
	values            map[string]*memValue
	locker            sync.RWMutex
}
type memValue struct {
	val    interface{}
	expire time.Time
	size   int64
}

func NewMemCache() memCache {
	return memCache{
		values: make(map[string]*memValue),
		locker: sync.RWMutex{},
	}
}
func (c *memCache) SetMaxMemory(size string) bool {
	var err error
	c.maxMemorySize, c.memorySizeStr, err = util.ParseStrSize(size)
	if err != nil {
		return false
	}
	log.Printf("初始化内存成功，大小为%dB", c.maxMemorySize)
	return true
}

func (c *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	c.locker.Lock()
	defer c.locker.Unlock()
	//确保map初始化
	if c.values == nil {
		c.values = make(map[string]*memValue)
	}
	v := &memValue{
		val:    val,
		expire: time.Now().Add(expire),
		size:   util.GetTypeSize(val),
	}
	if c.currMemorySize+v.size > c.maxMemorySize {
		log.Printf("缓存空间不足")
		return false
	}
	// 检查 key 是否已经存在
	if _, exists := c.values[key]; exists {
		log.Printf("该键值已存在，更新")
	}
	c.values[key] = v
	return true
}

func (c *memCache) Get(key string) (interface{}, bool) {
	c.locker.RLock()
	defer c.locker.RUnlock()
	if val, ok := c.values[key]; ok {
		if val.expire.After(time.Now()) {
			return val.val, true
		}
	}
	return nil, false
}

func (c *memCache) Del(key string) bool {
	c.locker.Lock()
	defer c.locker.Unlock()
	if val, found := c.values[key]; found {
		c.currMemorySize -= val.size
		delete(c.values, key)
		return true
	}
	return false
}

func (c *memCache) Exists(key string) bool {
	c.locker.Lock()
	defer c.locker.Unlock()
	_, found := c.values[key]
	return found
}

func (c *memCache) Flush() bool {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.values = make(map[string]*memValue)
	if len(c.values) == 0 {
		return true
	}
	return false
}

func (c *memCache) Keys() int64 {
	c.locker.RLock()
	defer c.locker.RUnlock()
	if c.values == nil {
		c.values = make(map[string]*memValue)
	}
	return int64(len(c.values))
}
