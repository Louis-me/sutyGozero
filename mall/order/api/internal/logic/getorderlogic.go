package logic

import (
	"context"
	"errors"
	"fmt"

	"study-gozero-demo/mall/order/api/internal/svc"
	"study-gozero-demo/mall/order/api/internal/types"
	"study-gozero-demo/mall/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	// 使用user rpc
	// 新增代码，注意这里的&user，应用的是user rpc的user/user.pb.go中内容,IdRequest其实就是对应入参（user.proto定义的）

	_, err = l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: int64(req.UserId),
	})
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	userInfo, err2 := l.svcCtx.OrderModel.FindOne(l.ctx, int64(req.Id))
	fmt.Println("userInfo=", userInfo, "error:", err2)
	if userInfo == nil {
		fmt.Println("iiiiiiii")
		return nil, errors.New("订单不存在")
	}
	return &types.OrderReply{
		Id:          req.Id,
		Productname: userInfo.Productname.String,
	}, nil

}
