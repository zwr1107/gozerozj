// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "gozerozj/user-api/internal/handler/user"
	"gozerozj/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.TestMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/add",
					Handler: user.AddUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/del",
					Handler: user.DelUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/edit",
					Handler: user.EditUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/info",
					Handler: user.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)
}
