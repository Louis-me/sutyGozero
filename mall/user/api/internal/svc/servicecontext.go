package svc

import (
	"study-gozero-demo/mall/user/api/internal/config"
	"study-gozero-demo/mall/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 引入model层的增删改查
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化mysql连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,
		// 调用user/model层代码
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
