package processes

import "testing"

func TestExecAndReturn(t *testing.T) {
	t.Log(Exec("‪C:/Go/bin/go.exe", "-v"))
}
