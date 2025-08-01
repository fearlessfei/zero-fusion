package subscribe

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"zero-fusion/app/demo/api/internal/logic/subscribe"
	"zero-fusion/app/demo/api/internal/svc"
	"zero-fusion/app/demo/api/internal/types"
	"zero-fusion/share/result"
)

// 订阅消息
func SubscribeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubscribeMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := subscribe.NewSubscribeLogic(r.Context(), svcCtx)
		resp, err := l.Subscribe(&req)
		result.HttpResult(r, w, resp, err)
	}
}
