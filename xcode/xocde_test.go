package xcode

import (
	"github.com/gogf/gf/g/test/gtest"
	"os"
	"testing"
)

func TestFindLatestDir(t *testing.T) {
	t.Log(Root)
	t.Log(findLatestDir(Root, "public"))
}

func TestTmpDir(t *testing.T) {
	t.Log(TmpDir())
	t.Log(TmpFile("test.json"))
}

func TestIsTesting(t *testing.T) {
	gtest.Assert(IsTesting(), true)
	t.Log(os.Args)
}
