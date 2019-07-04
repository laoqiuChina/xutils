package xip

import (
	"github.com/gogf/gf/g"
	"testing"
)

func TestXIP(t *testing.T) {
	ipFinder := GetInstance()
	g.Dump(ipFinder.Get("114.114.114.114"))
}
