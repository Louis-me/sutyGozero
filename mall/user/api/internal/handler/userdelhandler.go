package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"study-gozero-demo/mall/user/api/internal/logic"
	"study-gozero-demo/mall/user/api/internal/svc"
	"study-gozero-demo/mall/user/api/internal/types"
)

func UserDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserDelLogic(r.Context(), svcCtx)
		resp, err := l.UserDel(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
