package files

import (
	"git.qxtv1.com/st52/xutils/xcode"
	"testing"
)

func TestWriter_Write(t *testing.T) {
	tmpFile := NewFile(xcode.TmpFile("test.txt"))
	writer, err := tmpFile.Writer()
	if err != nil {
		t.Fatal(err)
	}

	//writer.Write([]byte("Hello,a"))
	//writer.Truncate()

	//writer.Seek(10)
	//writer.WriteString("ba")

	writer.Close()
}
