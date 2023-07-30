package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// 添加mysql配置
	Mysql struct {
		DataSource string
	}
	// 使用缓存
	CacheRedis cache.CacheConf
	// jwt的配置字段
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
