package errx

var (
	Success           = NewErrCodeMsg(200, "成功")
	ServerError       = NewErrCodeMsg(100001, "服务器繁忙，请稍后再试")
	DbError           = NewErrCodeMsg(100002, "数据库错误")
	RedisError        = NewErrCodeMsg(100003, "redis错误")
	RequestParamError = NewErrCodeMsg(100004, "请求参数错误")
)
