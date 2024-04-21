package main

import (
	"flag"
	"fmt"

	"GoZeroExample/application/user/rpc/internal/config"
	"GoZeroExample/application/user/rpc/internal/server"
	"GoZeroExample/application/user/rpc/internal/svc"
	"GoZeroExample/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/conf"
	service1 "github.com/zeromicro/go-zero/core/service"
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

		if c.Mode == service1.DevMode || c.Mode == service1.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
