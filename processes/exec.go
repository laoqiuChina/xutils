package processes

import (
	"io/ioutil"
	"os"
	"xutils/utils/string"
	"xutils/xcode"
)

func Exec(command string, args ...string) ([]byte, error) {
	randString := stringutil.Rand(32)
	tmpFile := xcode.TmpFile(randString + ".tmp")
	out, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return []byte{}, err
	}

	attrs := os.ProcAttr{
		Dir:   xcode.Root,
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, out, os.Stderr},
	}

	process, err := os.StartProcess(command, append([]string{command}, args...), &attrs)
	if err != nil {
		out.Close()
		_ = os.Remove(tmpFile)

		return []byte{}, err
	}
	_, err = process.Wait()
	if err != nil {
		out.Close()
		_ = os.Remove(tmpFile)
		return []byte{}, err
	}

	out.Close()
	outputData, err := ioutil.ReadFile(tmpFile)
	_ = os.Remove(tmpFile)

	if err != nil {
		return []byte{}, err
	}
	return outputData, nil
}

func ExecOut(out *os.File, command string, args ...string) error {
	attrs := os.ProcAttr{
		Dir:   xcode.Root,
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, out, os.Stderr},
	}

	process, err := os.StartProcess(command, append([]string{command}, args...), &attrs)
	if err != nil {
		return err
	}
	_, err = process.Wait()
	return err
}

func System32(file string) string {
	return os.Getenv("SystemRoot") + xcode.DS + "System32" + xcode.DS + file
}
