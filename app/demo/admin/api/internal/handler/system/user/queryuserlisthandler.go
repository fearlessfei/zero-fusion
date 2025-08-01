package user

import (
	"net/http"

	"zero-fusion/app/demo/admin/api/internal/logic/system/user"
	"zero-fusion/app/demo/admin/api/internal/svc"
	"zero-fusion/app/demo/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryUserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewQueryUserListLogic(r.Context(), svcCtx)
		resp, err := l.QueryUserList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
