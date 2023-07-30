package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"study-gozero-demo/mall/order/api/internal/logic"
	"study-gozero-demo/mall/order/api/internal/svc"
	"study-gozero-demo/mall/order/api/internal/types"
)

func orderAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderAddLogic(r.Context(), svcCtx)
		resp, err := l.OrderAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
