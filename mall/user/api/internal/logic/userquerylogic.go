package logic

import (
	"context"
	"errors"
	"fmt"

	"study-gozero-demo/mall/user/api/internal/svc"
	"study-gozero-demo/mall/user/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryLogic {
	return &UserQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserQueryLogic) UserQuery(req *types.UserQueryReq) (resp *types.UserQueryRes, err error) {
	// todo: add your logic here and delete this line

	fmt.Println("username", req.Username)
	res1, err := l.svcCtx.UserModel.FindAll(l.ctx, req.Username)
	fmt.Println("res1:", res1, "err:", err)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	var resList []*types.User
	for _, u := range res1 {
		var resVo types.User
		_ = copier.Copy(&resVo, u)
		resList = append(resList, &resVo)
	}
	return &types.UserQueryRes{UserList: resList}, nil

}
