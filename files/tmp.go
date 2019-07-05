package files

import "github.com/laoqiuChina/xutils/xcode"

func NewTmpFile(file string) *File {
	return NewFile(xcode.TmpFile(file))
}
