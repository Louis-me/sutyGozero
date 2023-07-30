package svc

import (
	"study-gozero-demo/mall/order/api/internal/config"
	"study-gozero-demo/mall/order/model"
	"study-gozero-demo/mall/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// 引入model层的增删改查
	OrderModel model.OrderModel
	// 引用最终生成user服务代码路径
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化mysql连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		// 调用order/model层代码
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		// 调用user rpc
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
