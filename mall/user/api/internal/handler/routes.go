// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"study-gozero-demo/mall/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/UserAdd",
				Handler: UserAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/UserDel",
				Handler: UserDelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/UserQuery",
				Handler: UserQueryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/UserEdit",
				Handler: UserEditHandler(serverCtx),
			},
		},
	)
}
