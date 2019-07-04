package ximg

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"image"
	"regexp"
)

type ImgServer struct {
	GetPath string
}

var regexpUrlParse *regexp.Regexp

var noImg *image.RGBA

func (m *ImgServer) Init() bool {
	var err error
	// 初始化正则表达式2
	regexpUrlParse, err = regexp.Compile("[a-z0-9]{32}")
	if err != nil {
		glog.Fatal("regexpUrlParse:", err)
		return false
	}
	// 创建 RGBA 画板大小 - 用于找不到图片时用
	noImg = image.NewRGBA(image.Rect(0, 0, 400, 400))
	return true
}

// Start 启动
func (m *ImgServer) Start() bool {
	if !m.Init() {
		return false
	}
	glog.Info("[ImgServer][params:" + m.String() + "]start... ")
	s := g.Server()

	// 缓存模式
	// 登录
	if m.GetPath == "" {
		glog.Error("[ImgServer]GetPath not set")
		return false
	}
	s.BindHandler(m.GetPath, m.ServeHTTP)
	return true
}

func (m ImgServer) ServeHTTP(r *ghttp.Request) {
	if r.Method == "GET" {
		m.Get(r.Response.ResponseWriter.ResponseWriter, r.Request)
		return
	}
	if r.Method == "POST" {
		m.Post(r.Response.ResponseWriter.ResponseWriter, r.Request)
		return
	}
}

// Start 结束
func (m *ImgServer) Stop() bool {
	glog.Info("[ImgServer]stop. ")
	return true
}

// String token解密方法
func (m *ImgServer) String() string {
	return ""
}
