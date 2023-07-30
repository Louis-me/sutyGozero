package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 添加mysql配置
	Mysql struct {
		DataSource string
	}
	// 使用缓存
	CacheRedis cache.CacheConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	// 引用user rpc
	UserRpc zrpc.RpcClientConf
}
