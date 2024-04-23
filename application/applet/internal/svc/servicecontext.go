package svc

import (
	"GoZeroExample/application/applet/internal/config"
	"GoZeroExample/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
