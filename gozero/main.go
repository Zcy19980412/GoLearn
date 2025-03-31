package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"gozero/intern/config"
)

var configFlag = flag.String("configFlag", "etc/config.yaml", "config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFlag, &c)

	var logConf = logx.LogConf{}
	conf.FillDefault(&logConf)
	logConf.Mode = "file"
	logc.MustSetup(logConf)
	defer logc.Close()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(writer http.ResponseWriter, request *http.Request) {
			logc.Info(context.Background(), "调用方法")
			writer.Write([]byte("hello world"))
		},
	})

	server.Start()

}
