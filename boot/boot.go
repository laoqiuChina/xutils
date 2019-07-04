package boot

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"xutils/ximg"
)

func init() {
	glog.Info("########service start...")

	v := g.View()
	c := g.Config()
	s := g.Server()

	path := ""
	// 配置对象及视图对象配置
	_ = c.AddPath(path + "config")

	v.SetDelimiters("${", "}")
	_ = v.AddPath(path + "template")

	// gLog配置
	logPath := c.GetString("log-path", "./logs")
	_ = glog.SetPath(logPath)
	glog.SetStdoutPrint(true)

	s.SetServerRoot("./public")
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	s.SetLogPath(logPath)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetPort(c.GetInt("http-port", 8080))

	glog.Info("########service finish.")
	imgServer := &ximg.ImgServer{
		SvPath: "/img",
	}
	imgServer.Start()
}
