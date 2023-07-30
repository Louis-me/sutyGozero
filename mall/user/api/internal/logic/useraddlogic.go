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

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserAddReq) (resp *types.UserAddRes, err error) {
	// todo: add your logic here and delete this line

	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	userInfo, err1 := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	fmt.Println("user_info", userInfo)
	// 如果查到已存在用户名，则不能新增
	if err1 == nil {
		return nil, errors.New(req.Username + "用户已经存在")
	}
	// 组装结构体
	var addReq model.User
	errCopy := copier.Copy(&addReq, req)
	if errCopy != nil {
		return nil, errCopy
	}
	// userInfoParam := model.User{}
	// userInfoParam.Pwd.String = req.Password
	// userInfoParam.Username.String = req.Username

	respd, err2 := l.svcCtx.UserModel.Insert(l.ctx, &addReq)
	fmt.Println("------respd:", respd, "error:", err2)
	if err2 != nil {
		fmt.Println("insert into error:", err)
		return nil, errors.New(err.Error())

	}

	// 登录成功返回数据
	return &types.UserAddRes{
		Name: req.Username,
	}, nil
}
