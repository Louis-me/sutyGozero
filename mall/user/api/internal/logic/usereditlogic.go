package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"study-gozero-demo/mall/user/api/internal/svc"
	"study-gozero-demo/mall/user/api/internal/types"
	"study-gozero-demo/mall/user/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserEditLogic {
	return &UserEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserEditLogic) UserEdit(req *types.UserEditReq) (resp *types.UserEditRes, err error) {
	// todo: add your logic here and delete this line

	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}
	if strings.TrimSpace(req.Username) == "admin" {
		return nil, errors.New("此用户不许修改")

	}

	userInfo, _ := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if userInfo == nil {
		return nil, errors.New("用户不存在")
	}
	// 组装结构体
	var editReq model.User
	errCopy := copier.Copy(&editReq, req)
	if errCopy != nil {
		return nil, errCopy
	}

	err2 := l.svcCtx.UserModel.Update(l.ctx, &editReq)
	if err2 != nil {
		fmt.Println("update  error:", err)
		return nil, errors.New("更新错误")

	}

	// 返回数据
	return &types.UserEditRes{
		Username: req.Username,
		Id:       req.Id,
	}, nil
}
