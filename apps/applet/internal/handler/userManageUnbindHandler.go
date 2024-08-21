package handler

import (
	"net/http"

	"application/apps/applet/internal/logic"
	"application/apps/applet/internal/svc"
	"application/apps/applet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserManageUnbindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserManageUnbindRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserManageUnbindLogic(r.Context(), svcCtx)
		resp, err := l.UserManageUnbind(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
