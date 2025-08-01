// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	loglogin "zero-fusion/app/demo/admin/api/internal/handler/log/login"
	logoperate "zero-fusion/app/demo/admin/api/internal/handler/log/operate"
	systemmenu "zero-fusion/app/demo/admin/api/internal/handler/system/menu"
	systemrole "zero-fusion/app/demo/admin/api/internal/handler/system/role"
	systemuser "zero-fusion/app/demo/admin/api/internal/handler/system/user"
	"zero-fusion/app/demo/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/deleteLoginLog",
					Handler: loglogin.DeleteLoginLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/queryLoginLogList",
					Handler: loglogin.QueryLoginLogListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/statisticsLoginLog",
					Handler: loglogin.StatisticsLoginLogHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/log/login"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/deleteOperateLog",
					Handler: logoperate.DeleteOperateLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/queryOperateLogList",
					Handler: logoperate.QueryOperateLogListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/log/operate"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addMenu",
					Handler: systemmenu.AddMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/deleteMenu",
					Handler: systemmenu.DeleteMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryMenuList",
					Handler: systemmenu.QueryMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateMenu",
					Handler: systemmenu.UpdateMenuHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/system/menu"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addRole",
					Handler: systemrole.AddRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/deleteRole",
					Handler: systemrole.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/queryRoleList",
					Handler: systemrole.QueryRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/queryRoleMenuList",
					Handler: systemrole.QueryRoleMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateRole",
					Handler: systemrole.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateRoleMenuList",
					Handler: systemrole.UpdateRoleMenuListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/system/role"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addUser",
					Handler: systemuser.AddUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/deleteUser",
					Handler: systemuser.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/queryUserList",
					Handler: systemuser.QueryUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryUserMenuList",
					Handler: systemuser.QueryUserMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/queryUserRoleList",
					Handler: systemuser.QueryUserRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateUser",
					Handler: systemuser.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateUserRoleList",
					Handler: systemuser.UpdateUserRoleListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/system/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: systemuser.UserLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/system"),
	)
}
