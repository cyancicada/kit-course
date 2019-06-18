package main

import (
	"context"
	"net/http"

	"github.com/yakaa/log4g"
	"ko/services"
)

func main() {

	// 1. 配置
	services.InitConfig()
	log4g.Init(log4g.Config{
		Path:   "logs",
		Stdout: true,
	})
	// 2. 日志系统
	httpLogger := log4g.StackLog

	// 3. 服务发现
	var ctx = context.Background()
	etcdClient := services.InitEtcd(ctx)
	registrar := services.RegisterSvc(etcdClient, httpLogger)
	// TODO: shutdown空指针报错
	defer registrar.Deregister()

	// 4. 路由服务
	var ucenterSvc services.UcenterServiceInterface
	ucenterSvc = services.UcenterService{}
	//ucenterSvc = middleware.InstrumentingMiddleware()(ucenterSvc)

	mux := http.NewServeMux()
	mux.Handle("/svc/ucenter/v1/", services.MakeHandler(ucenterSvc, httpLogger))
	mux.Handle("/svc/ucenter/v2/list/", services.MakeHandler(ucenterSvc, httpLogger))

	services.RunServer(mux, httpLogger, (*services.GetConfig())["server_port"])
}
