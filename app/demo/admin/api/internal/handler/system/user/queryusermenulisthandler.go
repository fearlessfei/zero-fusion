package user

import (
	"net/http"

	"zero-fusion/app/demo/admin/api/internal/logic/system/user"
	"zero-fusion/app/demo/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryUserMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewQueryUserMenuListLogic(r.Context(), svcCtx)
		resp, err := l.QueryUserMenuList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
