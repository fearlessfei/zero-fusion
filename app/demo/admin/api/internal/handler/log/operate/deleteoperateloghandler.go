package operate

import (
	"net/http"

	"zero-fusion/app/demo/admin/api/internal/logic/log/operate"
	"zero-fusion/app/demo/admin/api/internal/svc"
	"zero-fusion/app/demo/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteOperateLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteOperateLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := operate.NewDeleteOperateLogLogic(r.Context(), svcCtx)
		resp, err := l.DeleteOperateLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
