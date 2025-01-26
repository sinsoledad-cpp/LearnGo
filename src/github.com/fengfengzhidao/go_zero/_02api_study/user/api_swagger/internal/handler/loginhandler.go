package handler

import (
	"net/http"

	"api_swagger/internal/logic"
	"api_swagger/internal/svc"
	"api_swagger/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/fengfengzhidao/go-zero/common/response"
)

// 用户登录
func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)

	}
}
