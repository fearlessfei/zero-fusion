package role

import (
	"net/http"

	"zero-fusion/app/demo/admin/api/internal/logic/system/role"
	"zero-fusion/app/demo/admin/api/internal/svc"
	"zero-fusion/app/demo/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryRoleMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryRoleMenuListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewQueryRoleMenuListLogic(r.Context(), svcCtx)
		resp, err := l.QueryRoleMenuList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
