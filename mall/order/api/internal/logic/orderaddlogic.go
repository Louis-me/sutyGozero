package logic

import (
	"context"
	"errors"
	"fmt"

	"study-gozero-demo/mall/order/api/internal/svc"
	"study-gozero-demo/mall/order/api/internal/types"
	"study-gozero-demo/mall/order/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderAddLogic) OrderAdd(req *types.OrderAddReq) (resp *types.OrderAddRes, err error) {
	// todo: add your logic here and delete this line
	var addReq model.Order
	fmt.Println("req=", req)
	errCopy := copier.Copy(&addReq, req)
	if errCopy != nil {
		return nil, errCopy
	}

	respd, err2 := l.svcCtx.OrderModel.Insert(l.ctx, &addReq)
	fmt.Println("------respd:", respd, "error:", err2)
	if err2 != nil {
		fmt.Println("insert into error:", err)
		return nil, errors.New(err.Error())

	}

	// 登录成功返回数据
	return &types.OrderAddRes{
		Code:     1,
		Messsage: "success",
	}, nil

}
