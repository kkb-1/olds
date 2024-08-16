package handler

import (
	"net/http"

	"application/apps/applet/internal/logic"
	"application/apps/applet/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InviteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewInviteLogic(r.Context(), svcCtx)
		resp, err := l.Invite()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
