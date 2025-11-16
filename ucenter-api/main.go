// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
	"ucenter-api/internal/config"
	"ucenter-api/internal/handler"
	"ucenter-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()

	logx.MustSetup(logx.LogConf{Stat: false, Encoding: "plain"})

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	//先注册路由
	router := handler.NewRouters(server)
	//拿到路由后，再注册handler
	handler.RegisterHandlers(router, ctx)

	fmt.Printf("Starting server at ===================>>>>  %s:%d...\n", c.Host, c.Port)
	server.Start()
}
