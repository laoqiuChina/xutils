package files

import "xutils/xcode"

func NewTmpFile(file string) *File {
	return NewFile(xcode.TmpFile(file))
}
