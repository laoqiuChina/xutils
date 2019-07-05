package files

import "git.qxtv1.com/st52/xutils/xcode"

func NewTmpFile(file string) *File {
	return NewFile(xcode.TmpFile(file))
}
