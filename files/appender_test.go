package files

import (
	"github.com/laoqiuChina/xutils/xcode"
	"testing"
)

func Test_Appender(t *testing.T) {
	tmpFile := NewFile(xcode.TmpFile("test.txt"))
	appender, err := tmpFile.Appender()
	if err != nil {
		t.Fatal(err)
	}

	//appender.Lock()
	_, _ = appender.Append([]byte("Hello,a"))
	//appender.Truncate()

	_, _ = appender.AppendString("[ABC]")

	//appender.Unlock()
	t.Log(appender.Close())
}
