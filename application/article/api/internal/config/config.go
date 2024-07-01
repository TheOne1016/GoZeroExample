package config

import (
	"GoZeroExample/pkg/consul"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRPC    zrpc.RpcClientConf
	ArticleRPC zrpc.RpcClientConf
	Oss        struct {
		Endpoint         string
		AccessKeyId      string
		AccessKeySecret  string
		BucketName       string
		ConnectTimeout   int64 `json:",optional"`
		ReadWriteTimeout int64 `json:",optional"`
	}
	Consul consul.Conf
}
