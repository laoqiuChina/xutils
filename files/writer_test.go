package files

import (
	"github.com/laoqiuChina/xutils/xcode"
	"testing"
)

func TestWriter_Write(t *testing.T) {
	tmpFile := NewFile(xcode.TmpFile("te1st.txt"))
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
