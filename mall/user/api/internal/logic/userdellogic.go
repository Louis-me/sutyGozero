package logic

import (
	"context"
	"errors"
	"fmt"

	"study-gozero-demo/mall/user/api/internal/svc"
	"study-gozero-demo/mall/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDelLogic {
	return &UserDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDelLogic) UserDel(req *types.UserDelReq) (resp *types.UserDelRes, err error) {
	// todo: add your logic here and delete this line

	if req.Id == 23 {
		return nil, errors.New("此用户不可删除")
	}
	userInfo, _ := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if userInfo == nil {
		return nil, errors.New("用户不存在")
	}

	err2 := l.svcCtx.UserModel.Delete(l.ctx, req.Id)
	if err2 != nil {
		fmt.Println("del  error:", err)
		return nil, errors.New("删除错误")

	}

	// 返回数据
	return &types.UserDelRes{
		Username: userInfo.Username.String,
		Id:       userInfo.Id,
	}, nil
}
