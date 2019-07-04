package xstr

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

// 字符成长度
func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}
func MD5(str string) string {
	return MD5Bytes([]byte(str))
}
func MD5Bytes(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// 截取字符串
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}
func GetUUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	id := MD5(base64.URLEncoding.EncodeToString(b))
	return fmt.Sprintf("%s-%s-%s-%s-%s", id[0:8], id[8:12], id[12:16], id[16:20], id[20:])
}

// AbsoluteURL
// u : /1/2/3 ...
// baseURL : www.baidu.com
// Url : www.baidu.com/a/b/c
func AbsoluteURL(u string, baseURL, URL string) (string, error) {
	if strings.HasPrefix(u, "#") {
		return "", nil
	}

	var (
		_baseURL *url.URL
		_URL     *url.URL
	)

	_baseURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	if len(URL) > 0 {
		_URL, err = url.Parse(URL)
		if err != nil {
			return "", err
		}
	}

	var _base *url.URL
	if _baseURL != nil {
		_base = _baseURL
	} else {
		_base = _URL
	}

	if _base == nil {
		return u, nil
	}

	absURL, err := _base.Parse(u)
	if err != nil {
		return "", err
	}
	absURL.Fragment = ""
	if absURL.Scheme == "//" {
		absURL.Scheme = _URL.Scheme
	}
	return absURL.String(), nil
}
