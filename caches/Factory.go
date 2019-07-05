package caches

import (
	"git.qxtv1.com/st52/xutils/timers"
	"sync"
	"time"
)

// 操作类型
type CacheOperation = string

const (
	CacheOperationSet    = "set"
	CacheOperationDelete = "delete"
)

// 缓存管理器
type Factory struct {
	items       map[string]*Item
	maxSize     int64                               // @TODO 实现maxSize
	onOperation func(op CacheOperation, item *Item) // 操作回调
	locker      *sync.Mutex
	looper      *timers.Looper
}

// 创建一个新的缓存管理器
func NewFactory() *Factory {
	return newFactoryInterval(30 * time.Second)
}

func newFactoryInterval(duration time.Duration) *Factory {
	factory := &Factory{
		items:  map[string]*Item{},
		locker: &sync.Mutex{},
	}
	factory.looper = timers.Loop(duration, func(l *timers.Looper) {
		factory.clean()
	})
	return factory
}

// 设置缓存
func (c *Factory) Set(key string, value interface{}, duration ...time.Duration) *Item {
	item := new(Item)
	item.Key = key
	item.Value = value

	if len(duration) > 0 {
		item.expireTime = time.Now().Add(duration[0])
	} else {
		item.expireTime = time.Now().Add(3600 * time.Second)
	}
	c.locker.Lock()
	if c.onOperation != nil {
		_, ok := c.items[key]
		if ok {
			c.onOperation(CacheOperationDelete, item)
		}
	}
	c.items[key] = item
	c.locker.Unlock()
	if c.onOperation != nil {
		c.onOperation(CacheOperationSet, item)
	}
	return item
}

// 获取缓存
func (c *Factory) Get(key string) (value interface{}, found bool) {
	c.locker.Lock()
	defer c.locker.Unlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if item.IsExpired() {
		return nil, false
	}
	return item.Value, true
}

// 判断是否有缓存
func (c *Factory) Has(key string) bool {
	_, found := c.Get(key)
	return found
}

// 删除缓存
func (c *Factory) Delete(key string) {
	c.locker.Lock()
	item, ok := c.items[key]
	if ok {
		delete(c.items, key)
		if c.onOperation != nil {
			c.onOperation(CacheOperationDelete, item)
		}
	}
	c.locker.Unlock()
}

// 设置操作回调
func (c *Factory) OnOperation(f func(op CacheOperation, item *Item)) {
	c.onOperation = f
}

// 关闭
func (c *Factory) Close() {
	c.locker.Lock()
	defer c.locker.Unlock()
	if c.looper != nil {
		c.looper.Stop()
		c.looper = nil
	}
	c.items = map[string]*Item{}
}

// 重置状态
func (c *Factory) Reset() {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.items = map[string]*Item{}
}

// 清理过期的缓存
func (c *Factory) clean() {
	c.locker.Lock()
	defer c.locker.Unlock()
	for _, item := range c.items {
		if item.IsExpired() {
			delete(c.items, item.Key)
			if c.onOperation != nil {
				c.onOperation(CacheOperationDelete, item)
			}
		}
	}
}
