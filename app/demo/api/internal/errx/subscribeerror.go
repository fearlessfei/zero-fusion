package errx

import "zero-fusion/share/errx"

var (
	subscribeError = errx.NewErrCodeMsg(200100, "订阅失败")
)
