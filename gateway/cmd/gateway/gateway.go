package main

import (
	"context"

	"github.com/yakaa/log4g"
	"ko/gateway"
)


func main() {
	log4g.Init(log4g.Config{
		Path:   "logs",
		Stdout: true,
	})
	// 1. 配置


	// 2. 日志系统
	logger := log4g.StackLog

	// 3. 服务发现
	var ctx = context.Background()
	etcdClient := gateway.InitEtcd(ctx)

	// 1) 用户中心服务
	router := gateway.InitRouter(logger)
	router.Service("/svc/ucenter", etcdClient)
	///svc/ucenter/v1/
	router.Get("/svc/ucenter/v1/user/{param}")
	///svc/ucenter/v2/list
	router.Get("/svc/ucenter/v2/list")

	// 3) xx服务...

	// 4. 启动服务器
	gateway.RunServer(logger, gateway.Config["server_port"], router)
}
