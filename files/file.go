package files

import (
	"fmt"
	"github.com/gogf/gf/g/os/glog"
	stringutil "github.com/laoqiuChina/xutils/utils/string"
	"github.com/laoqiuChina/xutils/xcode"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 文件对象定义
type File struct {
	path string
}

// 包装新文件对象
func NewFile(path string) *File {
	return &File{
		path: path,
	}
}

// 取得文件名
func (f *File) Name() string {
	return filepath.Base(f.path)
}

// 取得文件统计信息
func (f *File) Stat() (*Stat, error) {
	stat, err := os.Stat(f.path)
	if err != nil {
		return nil, err
	}

	return &Stat{
		Name:    stat.Name(),
		Size:    stat.Size(),
		Mode:    stat.Mode(),
		ModTime: stat.ModTime(),
		IsDir:   stat.IsDir(),
	}, err
}

// 判断文件是否存在
func (f *File) Exists() bool {
	_, err := os.Stat(f.path)
	return !os.IsNotExist(err)
}

// 取得父级目录对象
func (f *File) Parent() *File {
	return NewFile(filepath.Dir(f.path))
}

// 判断是否为目录
func (f *File) IsDir() bool {
	stat, err := f.Stat()
	if err != nil {
		return false
	}

	return stat.IsDir
}

// 判断是否为文件
func (f *File) IsFile() bool {
	stat, err := f.Stat()
	if err != nil {
		return false
	}

	return stat.Mode.IsRegular()
}

// 读取最后更新时间
func (f *File) LastModified() (time.Time, error) {
	stat, err := f.Stat()
	if err != nil {
		return time.Unix(0, 0), err
	}
	return stat.ModTime, nil
}

// 读取文件尺寸
func (f *File) Size() (int64, error) {
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return stat.Size, nil
}

// 读取文件模式
func (f *File) Mode() (os.FileMode, error) {
	stat, err := f.Stat()
	if err != nil {
		return os.FileMode(0), err
	}
	return stat.Mode, nil
}

// 读取文件路径
func (f *File) Path() string {
	return f.path
}

// 读取文件绝对路径
func (f *File) AbsPath() (string, error) {
	p, err := filepath.Abs(f.path)
	if err != nil {
		return p, err
	}
	return p, nil
}

// 对文件内容进行Md5处理
func (f *File) Md5() (string, error) {
	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		return "", err
	}
	return stringutil.Md5(string(data)), err
}

// 读取文件内容
func (f *File) ReadAll() ([]byte, error) {
	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

// 读取文件内容并返回字符串形式
func (f *File) ReadAllString() (string, error) {
	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// 写入数据
func (f *File) Write(data []byte) error {
	writer, err := f.Writer()
	if err != nil {
		return err
	}
	writer.Lock()
	_ = writer.Truncate()
	_, err = writer.Write(data)
	_ = writer.Sync()
	writer.Unlock()
	_ = writer.Close()
	return err
}

// 写入字符串数据
func (f *File) WriteString(data string) error {
	return f.Write([]byte(data))
}

// 写入格式化的字符串数据
func (f *File) WriteFormat(format string, args ...interface{}) error {
	return f.WriteString(fmt.Sprintf(format, args...))
}

// 在文件末尾写入数据
func (f *File) Append(data []byte) error {
	appender, err := f.Appender()
	if err != nil {
		return err
	}
	appender.Lock()
	_, err = appender.Append(data)
	appender.Unlock()
	_ = appender.Close()
	return err
}

// 在文件末尾写入字符串数据
func (f *File) AppendString(data string) error {
	return f.Append([]byte(data))
}

// 取得文件扩展名，带点符号
func (f *File) Ext() string {
	return filepath.Ext(f.path)
}

// 目录下的文件
func (f *File) Child(filename string) *File {
	return NewFile(f.path + xcode.DS + filename)
}

// 列出目录下级的子文件对象
// 注意只会返回下一级，不会递归深入子目录
func (f *File) List() []*File {
	var result []*File

	if !f.IsDir() {
		return result
	}

	path, err := f.AbsPath()
	if err != nil {
		glog.Error(err)
		return result
	}

	fp, err := os.OpenFile(f.path, os.O_RDONLY, 0444)
	if err != nil {
		glog.Error(err)
		return result
	}

	defer fp.Close()
	names, err := fp.Readdirnames(-1)
	if err != nil {
		glog.Error(err)
		return result
	}

	for _, name := range names {
		result = append(result, NewFile(path+xcode.DS+name))
	}

	return result
}

// 使用模式匹配查找当前目录下的文件
func (f *File) Glob(pattern string) []*File {
	var result []*File
	matches, err := filepath.Glob(f.path + xcode.DS + pattern)
	if err != nil {
		return result
	}

	for _, path := range matches {
		result = append(result, NewFile(path))
	}

	return result
}

// 递归地对当前目录下的所有子文件、目录应用迭代器
func (f *File) Range(iterator func(file *File)) {
	if !f.Exists() || !f.IsDir() {
		return
	}

	for _, childFile := range f.List() {
		if childFile.IsDir() {
			childFile.Range(iterator)
		} else {
			iterator(childFile)
		}
	}
}

// 创建目录，但如果父级目录不存在，则会失败
func (f *File) Mkdir(perm ...os.FileMode) error {
	if len(perm) > 0 {
		return os.Mkdir(f.path, perm[0])
	}
	return os.Mkdir(f.path, 0777)
}

// 创建多级目录
func (f *File) MkdirAll(perm ...os.FileMode) error {
	if len(perm) > 0 {
		return os.MkdirAll(f.path, perm[0])
	}
	return os.MkdirAll(f.path, 0777)
}

// 创建文件
// 如果是目录，则使用 Mkdir() 和 MkdirAll()
func (f *File) Create() error {
	fp, err := os.Create(f.path)
	if err != nil {
		return err
	}
	fp.Close()
	return nil
}

// 修改文件的访问和修改时间为当前时间
func (f *File) Touch() error {
	now := time.Now()
	return os.Chtimes(f.path, now, now)
}

// 删除文件或目录，但如果目录不为空则会失败
func (f *File) Delete() error {
	if !f.IsDir() {
		return os.Remove(f.path)
	}
	return os.Remove(f.path)
}

// 判断文件或目录是否存在，然后删除文件或目录，如果目录不为空则会失败
func (f *File) DeleteIfExists() error {
	if !f.Exists() {
		return nil
	}
	if !f.IsDir() {
		return os.Remove(f.path)
	}
	return os.Remove(f.path)
}

// 拷贝文件
func (f *File) CopyTo(targetPath string) error {
	stat, err := f.Stat()
	if err != nil {
		return err
	}

	reader, err := os.OpenFile(f.path, os.O_RDONLY, 0444)
	if err != nil {
		return err
	}
	defer reader.Close()

	fp, err := os.OpenFile(targetPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, stat.Mode)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = io.Copy(fp, reader)
	return err
}

// 删除文件或目录，即使目录不为空也会删除
func (f *File) DeleteAll() error {
	if !f.IsDir() {
		return os.Remove(f.path)
	}
	return os.RemoveAll(f.path)
}

// 取得Writer
func (f *File) Writer() (*Writer, error) {
	writer := &Writer{}
	fp, err := os.OpenFile(f.path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	writer.file = fp
	writer.locker = &sync.Mutex{}
	return writer, nil
}

// 取得Appender
func (f *File) Appender() (*Appender, error) {
	appender := &Appender{}
	fp, err := os.OpenFile(f.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	appender.file = fp
	appender.locker = &sync.Mutex{}
	return appender, nil
}

// 取得Reader
func (f *File) Reader() (*Reader, error) {
	reader := &Reader{}
	fp, err := os.OpenFile(f.path, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}

	reader.file = fp
	return reader, nil
}

// 修改模式
func (f *File) Chmod(mode os.FileMode) error {
	return os.Chmod(f.path, mode)
}
