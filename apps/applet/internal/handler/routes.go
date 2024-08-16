// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"application/apps/applet/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/invite",
				Handler: InviteHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/userInfo",
				Handler: UpdateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/userInfo",
				Handler: GetUserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JWT.AccessSecret),
		rest.WithPrefix("/v1/user"),
	)
}
