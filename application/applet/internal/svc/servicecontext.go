package svc

import (
	"GoZeroExample/application/applet/internal/config"
	"GoZeroExample/application/user/rpc/user"
	"GoZeroExample/pkg/interceptors"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	//自定义拦截器
	userRpc := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	
	return &ServiceContext{
		Config:   c,
		UserRPC: user.NewUser(userRpc),
		BizRedis: redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
