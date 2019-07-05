package processes

import (
	"testing"
)

func TestNewProcess(t *testing.T) {
	process := NewProcess("‪C:\\Go\\bin\\go.exe", "-v")
	err := process.Start()
	if err != nil {
		t.Fatal("[ERROR]", err)
	}

	t.Log(process.Pid())

	//process.Wait()
}
