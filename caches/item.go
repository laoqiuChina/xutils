package caches

import "time"

// 缓存条目定义
type Item struct {
	Key        string
	Value      interface{}
	expireTime time.Time
}

// 设置值
func (i *Item) Set(value interface{}) *Item {
	i.Value = value
	return i
}

// 设置过期时间
func (i *Item) ExpireAt(expireTime time.Time) *Item {
	i.expireTime = expireTime
	return i
}

// 设置过期时长
func (i *Item) Expire(duration time.Duration) *Item {
	return i.ExpireAt(time.Now().Add(duration))
}

// 判断是否已过期
func (i *Item) IsExpired() bool {
	return time.Since(i.expireTime) > 0
}
