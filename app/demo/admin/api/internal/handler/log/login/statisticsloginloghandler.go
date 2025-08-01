package login

import (
	"net/http"

	"zero-fusion/app/demo/admin/api/internal/logic/log/login"
	"zero-fusion/app/demo/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func StatisticsLoginLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewStatisticsLoginLogLogic(r.Context(), svcCtx)
		resp, err := l.StatisticsLoginLog()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
