package main

import (
	"flag"
	"fmt"

	"GoZeroExample/application/user/rpc/internal/config"
	"GoZeroExample/application/user/rpc/internal/server"
	"GoZeroExample/application/user/rpc/internal/svc"
	"GoZeroExample/application/user/rpc/service"
	"GoZeroExample/pkg/consul"
	"GoZeroExample/pkg/interceptors"

	"github.com/zeromicro/go-zero/core/conf"
	cs "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == cs.DevMode || c.Mode == cs.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 自定义server端的拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	defer s.Stop()

	// 服务注册
	err := consul.Register(c.Consul, fmt.Sprintf("%s:%d", c.ServiceConf.Prometheus.Host, c.ServiceConf.Prometheus.Port))
	if err != nil {
		fmt.Printf("register consul error: %v\n", err)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
