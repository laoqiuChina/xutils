package xip

import (
	"github.com/gogf/gf/g"
	"testing"
)

func TestXIP(t *testing.T) {
	ipFinder := GetInstanceForPath("E:/go/gopath/src/xutils/data/ultimate.dat")
	g.Dump(ipFinder.Get("114.114.114.114"))
}
