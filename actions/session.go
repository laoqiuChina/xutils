package actions

import (
	"fmt"
	"github.com/gogf/gf/g/util/gconv"
)

// SESSION通用配置
type SessionConfig struct {
	Life   uint
	Secret string
}

// SESSION管理器接口
type SessionWrapper interface {
	Init(config *SessionConfig)
	Read(sid string) map[string]string
	WriteItem(sid string, key string, value string) bool
	Delete(sid string) bool
}

// Session定义
type Session struct {
	Sid     string
	Manager interface{}
}

// 设置sid
func (s *Session) SetSid(sid string) {
	s.Sid = sid
}

// 取得所有存储在session中的值
func (s *Session) Values() map[string]string {
	return s.Manager.(SessionWrapper).Read(s.Sid)
}

// 获取字符串值
func (s *Session) GetString(key string) string {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return ""
	}
	return value
}

// 获取int值
func (s *Session) GetInt(key string) int {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return 0
	}
	return gconv.Int(value)
}

// 获取int32值
func (s *Session) GetInt32(key string) int32 {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return 0
	}
	return gconv.Int32(value)
}

// 获取int64值
func (s *Session) GetInt64(key string) int64 {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return 0
	}
	return gconv.Int64(value)
}

// 获取uint值
func (s *Session) GetUint(key string) uint {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return 0
	}
	return gconv.Uint(value)
}

// 获取uint32值
func (s *Session) GetUint32(key string) uint32 {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return 0
	}
	return gconv.Uint32(value)
}

// 获取uint64值
func (s *Session) GetUint64(key string) uint64 {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return 0
	}
	return gconv.Uint64(value)
}

// 获取float32值
func (s *Session) GetFloat32(key string) float32 {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return 0
	}
	return gconv.Float32(value)
}

// 获取64值
func (s *Session) GetFloat64(key string) float64 {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return 0
	}
	return gconv.Float64(value)
}

// 获取bool值
func (s *Session) GetBool(key string) bool {
	values := s.Values()
	value, ok := values[key]
	if !ok {
		return false
	}
	return gconv.Bool(value)
}

// 写入值
func (s *Session) Write(key, value string) bool {
	return s.Manager.(SessionWrapper).WriteItem(s.Sid, key, value)
}

// 写入int值
func (s *Session) WriteInt(key string, value int) bool {
	return s.Write(key, fmt.Sprintf("%d", value))
}

// 写入int32值
func (s *Session) WriteInt32(key string, value int32) bool {
	return s.Write(key, fmt.Sprintf("%d", value))
}

// 写入int64值
func (s *Session) WriteInt64(key string, value int64) bool {
	return s.Write(key, fmt.Sprintf("%d", value))
}

// 写入uint值
func (s *Session) WriteUint(key string, value uint) bool {
	return s.Write(key, fmt.Sprintf("%d", value))
}

// 写入uint32值
func (s *Session) WriteUint32(key string, value uint32) bool {
	return s.Write(key, fmt.Sprintf("%d", value))
}

// 写入uint64值
func (s *Session) WriteUint64(key string, value uint64) bool {
	return s.Write(key, fmt.Sprintf("%d", value))
}

// 删除整个session
func (s *Session) Delete() bool {
	return s.Manager.(SessionWrapper).Delete(s.Sid)
}
