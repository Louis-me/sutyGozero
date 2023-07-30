package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"study-gozero-demo/mall/user/api/internal/logic"
	"study-gozero-demo/mall/user/api/internal/svc"
	"study-gozero-demo/mall/user/api/internal/types"
)

func UserAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserAddLogic(r.Context(), svcCtx)
		resp, err := l.UserAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
