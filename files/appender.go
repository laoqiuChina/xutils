package files

import (
	"os"
	"sync"
)

type Appender struct {
	file   *os.File
	locker *sync.Mutex
}

func NewAppender(path string) (*Appender, error) {
	return NewFile(path).Appender()
}

func (a *Appender) AppendString(s string) (n int, err error) {
	return a.file.WriteString(s)
}

func (a *Appender) Append(b []byte) (n int, err error) {
	return a.file.Write(b)
}

func (a *Appender) Truncate(size ...int64) error {
	if len(size) > 0 {
		return a.file.Truncate(size[0])
	}
	return a.file.Truncate(0)
}

func (a *Appender) Sync() error {
	return a.file.Sync()
}

func (a *Appender) Lock() {
	a.locker.Lock()
}

func (a *Appender) Unlock() {
	a.locker.Unlock()
}

func (a *Appender) Close() error {
	return a.file.Close()
}
