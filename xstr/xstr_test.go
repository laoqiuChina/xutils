package xstr

import (
	"git.qxtv1.com/st52/xutils/ximg"
	"testing"
)

func TestUrlParse(t *testing.T) {
	t.Log(ximg.UrlParse(MD5("chenyang")))
}

func TestGetAppPath(t *testing.T) {
	t.Log(GetAppPath())
}
func TestMD5(t *testing.T) {
	md5 := MD5("chenyang")
	t.Log(md5)
	if md5 == "b9a7256e85453ebbe6e133a9b490daef" {
		t.Log("TestMD5 done")
	} else {
		t.Error("md5 error : " + md5)
	}
}
func TestMD5Bytes(t *testing.T) {
	md5 := MD5Bytes([]byte("chenyang"))
	t.Log(md5)
	if md5 == "b9a7256e85453ebbe6e133a9b490daef" {
		t.Log("TestMD5 done")
	} else {
		t.Error("md5 error : " + md5)
	}
}
func TestRuneLen(t *testing.T) {
	md5 := MD5("chenyang")
	i := RuneLen(md5)
	if i != 32 {
		t.Fatal("Len error : " + md5)
	} else {
		t.Log("TestRuneLen done")
	}
}

func TestGetUUID(t *testing.T) {
	t.Log("TestGetUUID Begin")
	for i := 0; i < 100000; i++ {
		GetUUID()
	}
	t.Log("TestGetUUID End")
}

func TestAbsoluteURL(t *testing.T) {
	t.Log("TestAbsoluteURL Begin")
	if url, err := AbsoluteURL("/asda/asd", "www.xxoo.com", ""); err != nil {
		t.Error(err)
	} else {
		t.Log(url)
	}
	t.Log("TestAbsoluteURL End")
}

func TestSubstr(t *testing.T) {
	t.Log("TestSubstr Begin")
	md5 := MD5("chenyang")
	sub := Substr(md5, 0, 16)
	if sub == "b9a7256e85453ebb" {
		t.Log("TestSubstr done")
	} else {
		t.Error("md5 error : " + sub)
	}
	t.Log("TestSubstr End")
}
